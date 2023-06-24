package transporter

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQS struct {
	SQSClient *sqs.SQS
}

func (c SQS) HealthCheck(options *HealthCheckOptions) (bool, error) {
	_, err := c.SQSClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(options.QueueName),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// Publish module to publish sqs message for given input
func (c SQS) Publish(options *MessagePublishOptions) error {

	msgBody := aws.String(options.MessageBody.(string))

	queueUrl, queueUrlError := getSQSQueueURL(c, options.QueueName)
	if queueUrlError != nil {
		return queueUrlError
	}

	message := &sqs.SendMessageInput{
		MessageBody: msgBody,
		QueueUrl:    queueUrl,
	}

	if options.MessageGroupID != "" {
		message.MessageGroupId = aws.String(options.MessageGroupID)
	}

	_, err := c.SQSClient.SendMessage(message)
	if err != nil {
		logger.Errorf("error in publishing message for topic: %s, for: %s", options.TopicName, err)
		return err
	}

	return nil
}

// Consume module to receive sqs message concurrently with given channel and input
func (c SQS) Consume(options *MessageConsumeOptions) (*MessageReceiveResponse, error) {
	err := fmt.Errorf("method Consume is not implemented for SQS provider. Use BatchConsume with NumberOfMessages: 1")
	return nil, err
}

func (c SQS) BatchConsume(options *MessageConsumeOptions) ([]MessageReceiveResponse, error) {
	consumerTimeout := options.WaitTimeSeconds
	if consumerTimeout <= 0 {
		consumerTimeout = WaitTimeSeconds
	}

	queueUrl, queueUrlError := getSQSQueueURL(c, options.QueueName)
	if queueUrlError != nil {
		return nil, queueUrlError
	}

	result, err := c.SQSClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl,
		MaxNumberOfMessages: aws.Int64(int64(options.NumberOfMessages)),
		VisibilityTimeout:   aws.Int64(options.VisibilityTimeout),
		WaitTimeSeconds:     aws.Int64(options.WaitTimeSeconds),
	})

	if err != nil {
		logger.Errorf("error receive message from sqs, err: %s", err.Error())
		return nil, err
	}

	resp := make([]MessageReceiveResponse, len(result.Messages))

	for i, msg := range result.Messages {
		resp[i].MessageBody = *msg.Body
		resp[i].MessageReceiptHandle = *msg.ReceiptHandle
		resp[i].MessageID = *msg.MessageId
	}

	return resp, nil
}

func (c SQS) BatchPublish(options *MessagePublishOptions) error {
	return fmt.Errorf("method BatchPublish not implemented yet for SQS provider")
}

func (c SQS) DeleteMessage(queueName string, message *MessageReceiveResponse) error {
	queueUrl, queueUrlError := getSQSQueueURL(c, queueName)
	if queueUrlError != nil {
		return queueUrlError
	}
	_, err := c.SQSClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: aws.String(message.MessageReceiptHandle),
	})
	if err != nil {
		logger.Errorf("error in deleting SQS message with ID: %s, %s", message.MessageID, err)
		return err
	}
	logger.Info("deleted message on queue:", queueName)
	return nil
}

func (c SQS) BatchDeleteMessage(queueName string, messages []MessageReceiveResponse) error {
	return fmt.Errorf("method BatchDeleteMessage not implemented yet for SQS provider")
}

func getSQSQueueURL(c SQS, queueName string) (*string, error) {
	queueURL, err := c.SQSClient.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if err != nil || queueURL.QueueUrl == nil {
		logger.Errorf("error in generating queue URL for queueName, %s with error: %s", queueName, err)
		return nil, err
	}
	return queueURL.QueueUrl, nil

}
