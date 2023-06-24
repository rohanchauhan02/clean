package transporter

import (
	"fmt"
	"github.com/aliyun/aliyun-mns-go-sdk"
)

// MNS will implement the Client interface through pointer receivers
type MNS struct {
	Client ali_mns.MNSClient
}

func (mns *MNS) HealthCheck(options *HealthCheckOptions) (bool, error) {
	ali_mns.NewMNSQueue(options.QueueName, mns.Client)
	return true, nil
}

//Publish will push a message to the MNS queue with specified options and returns nil if no error
// TODO: should return a proper message response that abides to the Client interface
func (mns *MNS) Publish(options *MessagePublishOptions) error {
	msgBody, ok := options.MessageBody.(*string)
	if !ok {
		err := fmt.Errorf("message body is not of type string to publish to MNS queue: %s", options.QueueName)
		return err
	}

	queue := ali_mns.NewMNSQueue(options.QueueName, mns.Client)

	msg := ali_mns.MessageSendRequest{
		MessageBody:  *msgBody,
		Priority:     options.Priority,
		DelaySeconds: options.DelayInSeconds,
	}
	_, err := queue.SendMessage(msg)
	if err != nil {
		logger.Errorf("error in sending message to MNS queue: %s for: %s", options.QueueName, err)
		return err
	}
	return nil
}

//Consume will consume a single message from the specified endpoint which is initialized
// along with the client. Returns a valid MessageReceiveResponse if there is a message
// at the head of the queue and an error otherwise
func (mns *MNS) Consume(options *MessageConsumeOptions) (*MessageReceiveResponse, error) {

	consumerTimeout := options.WaitTimeSeconds
	if consumerTimeout <= 0 {
		consumerTimeout = WaitTimeSeconds
	}

	queue := ali_mns.NewMNSQueue(options.QueueName, mns.Client)

	responseChannel := make(chan ali_mns.MessageReceiveResponse)
	errChan := make(chan error)

	go queue.ReceiveMessage(responseChannel, errChan, consumerTimeout)
	for {
		select {
		case resp := <-responseChannel:
			{
				rs := &MessageReceiveResponse{
					MessageBody:              resp.MessageBody,
					MessageID:                resp.MessageId,
					MessageReceiptHandle:     resp.ReceiptHandle,
					MessageVisibilityTimeout: options.VisibilityTimeout,
				}
				if options.DeleteMessageAfterAck {
					_ = mns.DeleteMessage(options.QueueName, rs)
				}
				return rs, nil
			}
		case err := <-errChan:
			{
				logger.Errorf("error in consuming message from Alicloud MNS for queuename: %s for error: %s", options.QueueName, err)
				return nil, err
			}
		}
	}
}

//BatchPublish will push a max of 16 messages to a queue at a time with max of 64kb
//for all the messages, priority and delaySeconds will be the same across all messages
// TODO: should return a proper message response that abides to the Client interface
func (mns *MNS) BatchPublish(options *MessagePublishOptions) error {
	msgBodyArr, ok := options.MessageBody.([]string)
	if !ok {
		return fmt.Errorf("message body is not of type []string to batch publish to MNS queue: %s", options.QueueName)
	}

	queue := ali_mns.NewMNSQueue(options.QueueName, mns.Client)

	msgsRequest := make([]ali_mns.MessageSendRequest, len(msgBodyArr))
	for i, msgBody := range msgBodyArr {
		msgsRequest[i] = ali_mns.MessageSendRequest{
			MessageBody:  msgBody,
			Priority:     options.Priority,
			DelaySeconds: options.DelayInSeconds,
		}
	}
	_, err := queue.BatchSendMessage(msgsRequest...)
	if err != nil {
		logger.Errorf("error in sending message to MNS queue: %s for: %s", options.QueueName, err)
		return err
	}
	return nil
}

//BatchConsume consumes a max of 16 messages to a queue at a time, after the messages
//are received, they are in an inactive state. If DeleteMessageAfterAck is specified
//on consumer options, all messages that are consumed will be deleted via
//the BatchDelete routine
func (mns *MNS) BatchConsume(options *MessageConsumeOptions) ([]MessageReceiveResponse, error) {
	consumerTimeout := options.WaitTimeSeconds
	if consumerTimeout <= 0 {
		consumerTimeout = WaitTimeSeconds
	}

	queue := ali_mns.NewMNSQueue(options.QueueName, mns.Client)

	responseChannel := make(chan ali_mns.BatchMessageReceiveResponse, options.NumberOfMessages)
	errChan := make(chan error)

	go queue.BatchReceiveMessage(responseChannel, errChan, options.NumberOfMessages, consumerTimeout)
	for {
		select {
		case resp := <-responseChannel:
			{
				msgs := resp.Messages
				responses := make([]MessageReceiveResponse, len(msgs))
				for i, msg := range msgs {
					responses[i] = MessageReceiveResponse{
						MessageBody:              msg.MessageBody,
						MessageID:                msg.MessageId,
						MessageReceiptHandle:     msg.ReceiptHandle,
						MessageVisibilityTimeout: options.VisibilityTimeout,
					}
				}
				if options.DeleteMessageAfterAck {
					_ = mns.BatchDeleteMessage(options.QueueName, responses)
				}
				return responses, nil
			}
		case err := <-errChan:
			{
				logger.Errorf("error in consuming message from Alicloud MNS for queueName: %s for error: %s", options.QueueName, err)
				return nil, err
			}
		}
	}
}

//DeleteMessage will delete a message by taking the message's receiptHandle as the parameter
//deletion will happen only if timestamp of deletion is before NextVisibleTime of the message
//TODO: Add alarm/metric for failing in deleting messages
func (mns *MNS) DeleteMessage(queueName string, message *MessageReceiveResponse) error {
	queue := ali_mns.NewMNSQueue(queueName, mns.Client)

	ret, e := queue.ChangeMessageVisibility(message.MessageReceiptHandle, message.MessageVisibilityTimeout)
	if e != nil {
		logger.Errorf("error in changing message visibility MNS message with queue: %s, %s", queueName, e)
		return e
	}

	err := queue.DeleteMessage(ret.ReceiptHandle)
	if err != nil {
		logger.Errorf("error in deleting MNS message with queue: %s, %s", queueName, err)
		return err
	}
	logger.Info("Deleting message on queue:", queueName)
	return nil

}

//BatchDeleteMessage will delete a batch message by taking an array of message's receiptHandle as the parameter
//If there are some messages that cannot be deleted in the batch, the MNS API returns an array of FailedMessages
//by its receiptHandle.
//TODO: Add alarm/metric for failing in deleting messages
func (mns *MNS) BatchDeleteMessage(queueName string, messages []MessageReceiveResponse) error {

	queue := ali_mns.NewMNSQueue(queueName, mns.Client)

	receiptHandlers := make([]string, len(messages))
	for i, message := range messages {
		receiptHandlers[i] = message.MessageReceiptHandle
	}

	resp, err := queue.BatchDeleteMessage(receiptHandlers...)
	if err != nil {
		logger.Errorf("error in batch deleting message, %s", err)
		return err
	}
	for _, fm := range resp.FailedMessages {
		logger.Errorf("messages failed to be deleted with receiptHandle: %s, for error %s", fm.ReceiptHandle, fm.ErrorMessage)
	}
	return nil

}
