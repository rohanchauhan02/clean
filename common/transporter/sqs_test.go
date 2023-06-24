package transporter

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	awsAccessKey string
	awsSecretKey string
	sqsQueue     string
	sqsProvider  string
	sqsRegion    string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	awsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	sqsQueue = "qoala-common-test"
	sqsRegion = "ap-southeast-1"
	sqsProvider = "aws"
}

func TestNewSQS(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	t.Run("test ok new SQS client initialization", func(t *testing.T) {
		_, err := NewClient(conf)
		assert.Nil(t, err, "New SQS client is successfully created")
	})
}

func TestNewClientInvalidProvider(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  "INVALID_VENDOR",
		Region:    sqsRegion,
	}
	t.Run("test new SQS client initalization expect error", func(t *testing.T) {
		_, err := NewClient(conf)
		assert.Error(t, err)
	})
}

func TestPublishSQSWithAWSRole(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)
	options := &MessagePublishOptions{
		QueueName:   sqsQueue,
		MessageBody: aws.String("message from wrapper"),
	}

	t.Run("test SQS publish with aws role", func(t *testing.T) {
		err := client.Publish(options)
		assert.Nil(t, err, "ok publish sqs message")
	})
}

func TestPublishSQSWithAWSIAM(t *testing.T) {
	conf := &ClientOptions{
		QueueName:       sqsQueue,
		Provider:        sqsProvider,
		Region:          sqsRegion,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
	}

	client, _ := NewClient(conf)
	options := &MessagePublishOptions{
		QueueName:   sqsQueue,
		MessageBody: aws.String("message from wrapper"),
	}

	t.Run("test SQS publish with aws iam", func(t *testing.T) {
		err := client.Publish(options)
		assert.Nil(t, err, "ok publish sqs message")
	})
}

func TestPublishSQSWithAWSRoleInvalidBody(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)
	options := &MessagePublishOptions{
		QueueName:   sqsQueue,
		MessageBody: nil,
	}

	t.Run("test SQS publish with aws iam invalid body", func(t *testing.T) {
		err := client.Publish(options)
		assert.Error(t, err)
	})
}

func TestPublishSQSWithAWSIAMInvalidBody(t *testing.T) {
	conf := &ClientOptions{
		QueueName:       sqsQueue,
		Provider:        sqsProvider,
		Region:          sqsRegion,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
	}

	client, _ := NewClient(conf)
	options := &MessagePublishOptions{
		QueueName:   sqsQueue,
		MessageBody: nil,
	}

	t.Run("test SQS publish with aws iam invalid body", func(t *testing.T) {
		err := client.Publish(options)
		assert.Error(t, err)
	})
}

func TestConsumeSQSWithAWSIAM(t *testing.T) {
	conf := &ClientOptions{
		QueueName:       sqsQueue,
		Provider:        sqsProvider,
		Region:          sqsRegion,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
	}

	client, _ := NewClient(conf)

	options := &MessageConsumeOptions{
		QueueName:         sqsQueue,
		VisibilityTimeout: 10,
		NumberOfMessages:  4,
		WaitTimeSeconds:   3,
	}

	t.Run("test SQS consume with valid options", func(t *testing.T) {
		resp, err := client.BatchConsume(options)
		assert.Nil(t, err)
		assert.True(t, len(resp) <= 4)
		assert.NotEmpty(t, resp[0].MessageBody)
		assert.NotEmpty(t, resp[0].MessageID)
		assert.NotEmpty(t, resp[0].MessageReceiptHandle)
	})
}

func TestConsumeSQSWithAWSRole(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)

	options := &MessageConsumeOptions{
		QueueName:         sqsQueue,
		VisibilityTimeout: 10,
		NumberOfMessages:  4,
		WaitTimeSeconds:   3,
	}

	t.Run("test SQS role consume with valid options", func(t *testing.T) {
		resp, err := client.BatchConsume(options)
		assert.Nil(t, err)
		assert.True(t, len(resp) <= 4)
		assert.NotEmpty(t, resp[0].MessageBody)
		assert.NotEmpty(t, resp[0].MessageID)
		assert.NotEmpty(t, resp[0].MessageReceiptHandle)
	})
}

func TestSQSPublishJSONMessage(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)

	m := map[string]int{"id": 1, "quantity": 10}
	jsonString, _ := json.Marshal(m)

	options := &MessagePublishOptions{
		QueueName:   sqsQueue,
		MessageBody: string(jsonString),
	}

	t.Run("test SQS publish json message", func(t *testing.T) {
		err := client.Publish(options)
		assert.Error(t, err)
	})
}

func TestSQSDeleteMessageWithInvalidReceiptHandle(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)

	response := &MessageReceiveResponse{
		MessageBody:          aws.String("'{key': message from wrapper}"),
		MessageReceiptHandle: "INVALID_RECEIPT",
	}

	t.Run("test SQS delete message with invalid receipt handle", func(t *testing.T) {
		err := client.DeleteMessage(queueName, response)
		assert.Error(t, err)
	})
}

func TestSQSDeleteMessageRoutine(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)

	m := map[string]int{"id": 1, "quantity": 10}
	jsonByte, _ := json.Marshal(m)
	jsonString := string(jsonByte)

	publishOptions := &MessagePublishOptions{
		QueueName:   sqsQueue,
		MessageBody: &jsonString,
	}

	err := client.Publish(publishOptions)
	if err != nil {
		assert.Fail(t, "error in publishing message in sqs delete message routine")
	}

	consumerOptions := &MessageConsumeOptions{
		DeleteMessageAfterAck: false,
		NumberOfMessages:      1,
		QueueName:             sqsQueue,
	}

	msgResponse, err := client.BatchConsume(consumerOptions)
	if err != nil {
		assert.Fail(t, "error in consuming message in sqs delete message routine")
	}

	t.Run("test SQS delete message routine", func(t *testing.T) {
		err := client.DeleteMessage(sqsQueue, &msgResponse[0])
		assert.Nil(t, err)
	})
}

func TestSQSPublishWithMessageGroupID(t *testing.T) {
	conf := &ClientOptions{
		QueueName: sqsQueue,
		Provider:  sqsProvider,
		Region:    sqsRegion,
	}

	client, _ := NewClient(conf)

	m := map[string]int{"id": 1, "quantity": 10}
	jsonString, _ := json.Marshal(m)

	options := &MessagePublishOptions{
		QueueName:      sqsQueue,
		MessageBody:    string(jsonString),
		MessageGroupID: "test-group-id",
	}

	t.Run("test SQS publish message group id", func(t *testing.T) {
		err := client.Publish(options)
		assert.Error(t, err)
	})
}
