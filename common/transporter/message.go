package transporter

//MessagePublishOptions used when publishing a message across different
//vendor implementors.
//* MessageBody will be the payload sent to the queue
//* Alicloud casts the MessageBody to String
//* NATS casts the MessageBody to []byte
//* TopicName is the name of topic - this is optional for Alicloud
//* Priority is used in a PriorityQueue setting in Alicloud - (only used for Alicloud)
//* value for Priority should be between 1 to 16
//* DelayInSeconds states the messages cannot be consumed until the period specified by the DelayInSeconds parameter ends.
type MessagePublishOptions struct {
	MessageBody    interface{} `json:"message_body"`
	QueueName      string
	TopicName      string
	Priority       int64
	DelayInSeconds int64
	MessageGroupID string
}

//MessageConsumeOptions used when consuming a message across different
//vendor implementors.
//* TopicName is the name of topic - this is optional for Alicloud
//* DeleteMessageAfterAck - this will delete/ack the message after being consumed
//* VisibilityTimeout - will be set to 30 seconds by default and sets the timeout to which the consumer needs to return
// before returning.
//* NumberOfMessages - this is used when consuming a batch of messages, this is used to specify the max number of messages
//* that can be received in a batch
type MessageConsumeOptions struct {
	QueueName             string
	TopicName             string
	DeleteMessageAfterAck bool
	VisibilityTimeout     int64
	NumberOfMessages      int32
	WaitTimeSeconds       int64
}

//MessageReceiveResponse is a wrapper for the consumer response from
//different vendor implementors
//* MessageSubject - similar to nats.Msg subject
//* MessageBody - payload from consumer
//* MessageId - unique ID when message is sent
type MessageReceiveResponse struct {
	MessageBody              interface{}
	MessageSubject           string
	MessageID                string
	MessageReceiptHandle     string
	MessageVisibilityTimeout int64
}
