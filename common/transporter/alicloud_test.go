package transporter

import (
	"os"

	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	mnsEndpoint 		string
	accessKey   		string
	secretKey   		string
	queueName   		string
	region      		string
	alicloudProvider	string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mnsEndpoint = os.Getenv("ALICLOUD_MNS_ENDPOINT")
	accessKey = os.Getenv("ALICLOUD_ACCESS_KEY")
	secretKey = os.Getenv("ALICLOUD_SECRET_KEY")
	queueName = os.Getenv("ALICLOUD_QUEUE_NAME")
	alicloudProvider = "alicloud"

}

func TestNewMNSClient(t *testing.T) {

	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
		QueueName:       queueName,
	}

	t.Run("test new MNS client initialization", func(t *testing.T) {
		_, err := NewClient(&conf)
		assert.Nil(t, err, "New MNS client is successfully created")
	})
}

func TestNewMNSClientWithInvalidProvider(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        "INVALID_VENDOR",
	}

	t.Run("test new MNS client initialization", func(t *testing.T) {
		_, err := NewClient(&conf)
		assert.Error(t, err)
	})
}

func TestAlicloudPublishWithNoPriority(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessagePublishOptions{
		MessageBody: "message from wrapper",
		QueueName:   queueName,
	}

	t.Run("test Alicloud publish with no priority in MessageOptions", func(t *testing.T) {
		err := client.Publish(options)
		assert.Error(t, err)
	})
}

func TestAlicloudPublishWithInvalidMessageBody(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessagePublishOptions{
		MessageBody: nil,
		Priority:    3,
		QueueName:   queueName,
	}

	t.Run("test Alicloud publish with invalid message body", func(t *testing.T) {
		err := client.Publish(options)
		assert.Error(t, err)
	})
}

func TestAlicloudPublishWithValidOptions(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	msgBody := "message from wrapper"
	options := &MessagePublishOptions{
		MessageBody: &msgBody,
		Priority:    8,
		QueueName:   queueName,
	}

	t.Run("test Alicloud publish with valid MessageOptions", func(t *testing.T) {
		err := client.Publish(options)
		assert.Nil(t, err)
	})
}

func TestAlicloudConsumeWithNegativeConsumerTimeout(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessageConsumeOptions{
		WaitTimeSeconds: -1,
		QueueName:       queueName,
	}

	t.Run("test Alicloud consume with negative consumer timeout", func(t *testing.T) {
		resp, err := client.Consume(options)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp.MessageBody)
		assert.NotEmpty(t, resp.MessageID)
		assert.NotEmpty(t, resp.MessageReceiptHandle)
	})
}

func TestAliCloudDeleteMessageAfterConsume(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	msgBody := "message from wrapper"
	publishOptions := &MessagePublishOptions{
		MessageBody:    &msgBody,
		DelayInSeconds: 600,
		Priority:       16,
		QueueName:      queueName,
	}

	err := client.Publish(publishOptions)
	if err != nil {
		assert.Fail(t, "Publish message failed on testing delete message after consume")
	}

	consumerOptions := &MessageConsumeOptions{
		DeleteMessageAfterAck: true,
		QueueName:             queueName,
		VisibilityTimeout:     30,
	}

	t.Run("test Alicloud delete message after consuming message", func(t *testing.T) {
		resp, err := client.Consume(consumerOptions)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp.MessageBody)
		assert.NotEmpty(t, resp.MessageID)
		assert.NotEmpty(t, resp.MessageReceiptHandle)
		assert.NotEmpty(t, resp.MessageVisibilityTimeout)
	})
}

func TestAlicloudBatchPublishWithInvalidMessageBody(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessagePublishOptions{
		MessageBody: []int{1, 2},
		Priority:    8,
		QueueName:   queueName,
	}

	t.Run("test Alicloud batch publish with invalid MessageBody", func(t *testing.T) {
		err := client.BatchPublish(options)
		assert.Error(t, err)
	})
}

func TestAlicloudBatchPublishWithNoPriority(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessagePublishOptions{
		MessageBody: []string{"h1", "h2", "h3"},
		QueueName:   queueName,
	}

	t.Run("test Alicloud batch publish with no priority", func(t *testing.T) {
		err := client.BatchPublish(options)
		assert.Error(t, err)
	})
}

func TestAlicloudBatchPublishWithGreaterThan16Messages(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	strArray := make([]string, 20)
	for i := range strArray {
		strArray[i] = "val"
	}
	options := &MessagePublishOptions{
		MessageBody: strArray,
		QueueName:   queueName,
	}

	t.Run("test Alicloud request publish greather than 16 messages", func(t *testing.T) {
		err := client.BatchPublish(options)
		assert.Error(t, err)
	})
}

func TestAlicloudBatchPublishWithValidMessageOptions(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessagePublishOptions{
		MessageBody: []string{"h1", "h2", "h3"},
		Priority:    8,
		QueueName:   queueName,
	}

	t.Run("test Alicloud batch publish with valid MessageOptions", func(t *testing.T) {
		err := client.BatchPublish(options)
		assert.Nil(t, err)
	})
}
func TestAlicloudBatchConsumeWithNegativeConsumerTimeout(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessageConsumeOptions{
		VisibilityTimeout: -1,
		QueueName:         queueName,
	}

	t.Run("test Alicloud consume with negative consumer timeout", func(t *testing.T) {
		resp, err := client.BatchConsume(options)
		assert.Nil(t, err)
		assert.NotNil(t, resp[0].MessageBody)
		assert.NotEmpty(t, resp[0].MessageReceiptHandle)
	})
}
func TestAlicloudBatchConsumeWithGreaterThan16Messages(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessageConsumeOptions{
		VisibilityTimeout: -1,
		NumberOfMessages:  30,
		QueueName:         queueName,
	}

	t.Run("test Alicloud request consume greather than 16 messages", func(t *testing.T) {
		_, err := client.BatchConsume(options)
		assert.Error(t, err)
	})
}

func TestAliCloudDeleteMessageAfterBatchConsume(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	publishOptions := &MessagePublishOptions{
		MessageBody:    []string{"h1", "h2"},
		DelayInSeconds: 600,
		Priority:       16,
		QueueName:      queueName,
	}

	err := client.BatchPublish(publishOptions)
	if err != nil {
		assert.Fail(t, "batch publish message failed on testing delete message after consume")
	}

	consumerOptions := &MessageConsumeOptions{
		DeleteMessageAfterAck: true,
		NumberOfMessages:      2,
		QueueName:             queueName,
	}

	t.Run("test Alicloud delete message after consuming message", func(t *testing.T) {
		_, err := client.BatchConsume(consumerOptions)
		assert.Nil(t, err)
	})
}

func TestAlicloudDeleteMessageWithInvalidReceiptHandle(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	options := &MessageReceiveResponse{
		MessageBody:          "SOME_MESSAGE_BODY",
		MessageReceiptHandle: "INVALID_RECEIPT",
	}

	t.Run("test Alicloud delete message with invalid receipt handle", func(t *testing.T) {
		err := client.DeleteMessage(queueName, options)
		assert.Error(t, err)
	})
}

func TestAlicloudDeleteMessageRoutine(t *testing.T) {
	conf := ClientOptions{
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
		Endpoint:        mnsEndpoint,
		Provider:        alicloudProvider,
	}

	client, _ := NewClient(&conf)

	msgBody := "message from wrapper"
	publishOptions := &MessagePublishOptions{
		MessageBody:    &msgBody,
		DelayInSeconds: 600,
		Priority:       16,
		QueueName:      queueName,
	}

	err := client.Publish(publishOptions)
	if err != nil {
		assert.Fail(t, "publish message failed on testing delete message routine")
	}

	consumerOptions := &MessageConsumeOptions{
		DeleteMessageAfterAck: false,
		QueueName:             queueName,
		VisibilityTimeout:     30,
	}

	msgResponse, err := client.Consume(consumerOptions)
	if err != nil {
		assert.Fail(t, "consume message failed on testing delete message routine")
	}

	t.Run("test Alicloud call DeleteMessage", func(t *testing.T) {
		assert.Equal(t, consumerOptions.VisibilityTimeout, msgResponse.MessageVisibilityTimeout)
		err := client.DeleteMessage(queueName, msgResponse)
		assert.Nil(t, err)
	})
}
