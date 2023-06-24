package transporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNewNats(t *testing.T) {
//         conf := ClientOptions{
//                 Host: "localhost",
//                 Port: "4222",
//         }
//
//         t.Run("test ok factory method nats", func(t *testing.T) {
//                 respIntf := NewNats(conf)
//                 assert.NotNil(t, respIntf, "test ok not nil")
//         })
// }

func TestOpenNatsConn(t *testing.T) {
	conf := ClientOptions{
		Host:     "localhost",
		Port:     "4222",
		Provider: "nats",
	}

	t.Run("test ok open nats conn", func(t *testing.T) {
		_, err := NewClient(&conf)
		assert.Nil(t, err, "test ok err nil")
	})
}

func TestConsumerNats(t *testing.T) {
	conf := ClientOptions{
		Host:     "localhost",
		Port:     "4222",
		Provider: "nats",
	}

	options := MessageConsumeOptions{
		TopicName: "TEST-TOPIC",
	}

	var natsCl Client
	t.Run("test ok open nats conn", func(t *testing.T) {
		cl, err := NewClient(&conf)
		natsCl = cl
		assert.Nil(t, err, "test ok err nil")
	})

	t.Run("test ok consume nats message", func(t *testing.T) {
		msg, err := natsCl.Consume(&options)
		assert.Nil(t, err)
		assert.Empty(t, msg.MessageBody, "test ok consume nats message")
	})

}

func TestPublishNats(t *testing.T) {
	conf := ClientOptions{
		Host:     "localhost",
		Port:     "4222",
		Provider: "nats",
	}

	publisherOptions := MessagePublishOptions{
		MessageBody: "HALO",
		TopicName:   "TEST-TOPIC",
	}

	consumerOptions := MessageConsumeOptions{
		TopicName: "TEST-TOPIC",
	}

	var natsCl Client
	t.Run("test ok open nats conn", func(t *testing.T) {
		cl, err := NewClient(&conf)
		natsCl = cl
		assert.Nil(t, err, "test ok err nil")
	})

	t.Run("test ok publish nats message", func(t *testing.T) {
		err := natsCl.Publish(&publisherOptions)
		assert.Nil(t, err, "test ok publish no error")
	})

	t.Run("test ok consume nats message", func(t *testing.T) {
		msg, err := natsCl.Consume(&consumerOptions)
		assert.Nil(t, err)
		msgBody := msg.MessageBody.([]byte)
		assert.NotNil(t, string(msgBody), "test ok consume nats message")
	})

}
