package transporter

import (
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
)

// NatsClient will implement the Client interface through pointer receivers
type NatsClient struct {
	EncodedConnection *nats.EncodedConn
}

func (natsClient NatsClient) HealthCheck(options *HealthCheckOptions) (bool, error) {
	panic("Health check not implemented yet")
}

// Consume module to consume nats message for given topic
func (natsClient NatsClient) Consume(options *MessageConsumeOptions) (*MessageReceiveResponse, error) {

	natsMsg := &nats.Msg{}
	_, err := natsClient.EncodedConnection.QueueSubscribe(options.TopicName, strings.ToLower(options.TopicName), func(message *nats.Msg) {
		logger.Info("Consumed mesage from NATS")
		logger.Info(string(message.Data))
		natsMsg = message
	})

	if err != nil {
		return nil, err
	}

	return &MessageReceiveResponse{
		MessageBody:    natsMsg.Data,
		MessageSubject: natsMsg.Subject,
	}, nil
}
func (natsClient NatsClient) BatchConsume(options *MessageConsumeOptions) ([]MessageReceiveResponse, error) {
	return nil, fmt.Errorf("Method BatchConsume not implemented for NATS provider")
}

// Publish module to publish nats message for given topic and data
func (natsClient NatsClient) Publish(options *MessagePublishOptions) error {

	err := natsClient.EncodedConnection.Publish(options.TopicName, options.MessageBody)
	if err != nil {
		logger.Errorf("Error in publishing message for topic: %s, for: %s", options.TopicName, err)
		return err
	}
	return nil
}

func (natsClient NatsClient) BatchPublish(options *MessagePublishOptions) error {
	return fmt.Errorf("Method BatchPublish not implemented yet for NATS provider")
}

func (natsClient NatsClient) DeleteMessage(queueName string, message *MessageReceiveResponse) error {
	return fmt.Errorf("Method DeleteMessage not implemented yet for NATS provider")
}

func (natsClient NatsClient) BatchDeleteMessage(queueName string, messages []MessageReceiveResponse) error {
	return fmt.Errorf("Method BatchDeleteMessage not implemented yet for NATS provider")
}
