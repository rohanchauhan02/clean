package transporter

import (
	"fmt"
	"strings"

	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/nats-io/nats.go"
)

//ClientOptions is a shared options used across different
//vendor implementations to initialize the client used for
//SQS, MNS, NATS etc.
type ClientOptions struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	Provider        string `json:"provider"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Region          string `json:"region"`
	QueueName       string `json:"queueName"`
}

type HealthCheckOptions struct {
	QueueName string `json:"queueName"`
}

//Client is interface for vendor provider wrapper to implement
//publish and consume methods from the cloud provider.
//Options are shared across vendors so some of the Publish/Consume options
//may be unused across different implementors
type Client interface {
	HealthCheck(options *HealthCheckOptions) (bool, error)
	Publish(options *MessagePublishOptions) error
	BatchPublish(options *MessagePublishOptions) error
	Consume(options *MessageConsumeOptions) (*MessageReceiveResponse, error)
	BatchConsume(options *MessageConsumeOptions) ([]MessageReceiveResponse, error)
	DeleteMessage(queueName string, message *MessageReceiveResponse) error
	BatchDeleteMessage(queueName string, messages []MessageReceiveResponse) error
}

//NewClient initializes a new client depending on the provided/vendor type
//The client will be used in the vendor *.go file to access the SDK queue
//methods
func NewClient(options *ClientOptions) (Client, error) {
	provider := strings.ToLower(options.Provider)
	switch provider {
	case "alicloud":
		{

			client := ali_mns.NewAliMNSClient(
				options.Endpoint,
				options.AccessKeyID,
				options.AccessKeySecret)

			mns := &MNS{
				Client: client,
			}
			return mns, nil
		}
	case "nats":
		{
			connString := fmt.Sprintf("%s:%s", options.Host, options.Port)
			client, err := nats.Connect(connString)
			if err == nil {
				encodedConn, err := nats.NewEncodedConn(client, nats.JSON_ENCODER)
				if err == nil {
					natsClient := &NatsClient{
						EncodedConnection: encodedConn,
					}
					return natsClient, nil
				}
			}
			logger.Error("Error in initializing NATS client")
			return nil, err
		}
	case "aws":
		{
			awsSess, err := GetAWSSession(options.AccessKeyID, options.AccessKeySecret, options.Region)
			if err == nil {
				client := sqs.New(awsSess)
				sqsClient := &SQS{
					SQSClient: client,
				}
				return sqsClient, nil
			}

			return nil, err
		}
	}
	return nil, fmt.Errorf("provider:\"%s\" is not supported in transporter module", options.Provider)
}