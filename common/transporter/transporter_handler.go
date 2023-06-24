package transporter

import "github.com/pkg/errors"

type (
	ServiceConsumer interface {
		Consume(channelMessages chan *MessageReceiveResponse)
		GetWorkerPool() int
		Acknowledge(message *MessageReceiveResponse) error
	}

	consumer struct {
		transporterClient Client
		queueName         string
		workerPool        int
		numberOfMessage   int
		waitTimeSecond    int
		visibilityTimeout int
	}

	ConsumerOption struct {
		TransporterClient Client
		QueueName         string
		WorkerPool        int
		NumberOfMessage   int
		WaitTimeSecond    int
		VisibilityTimeout int
	}
)

func NewServiceConsumer(option *ConsumerOption) (ServiceConsumer, error) {
	if option == nil {
		return nil, errors.New("option is empty")
	}

	return &consumer{
		transporterClient: option.TransporterClient,
		queueName:         option.QueueName,
		workerPool:        option.WorkerPool,
		numberOfMessage:   option.NumberOfMessage,
		waitTimeSecond:    option.WaitTimeSecond,
		visibilityTimeout: option.VisibilityTimeout,
	}, nil
}

func (c *consumer) Consume(channelMessages chan *MessageReceiveResponse) {
	for {
		result, err := c.transporterClient.BatchConsume(&MessageConsumeOptions{
			QueueName:         c.queueName,
			VisibilityTimeout: int64(c.visibilityTimeout),
			NumberOfMessages:  int32(c.numberOfMessage),
			WaitTimeSeconds:   int64(c.waitTimeSecond),
		})

		if err != nil {
			logger.Error(errors.Wrap(err, "failed to retrieve message"))
			continue
		}

		for _, receivedMessage := range result {
			channelMessages <- &receivedMessage
		}
	}
}

func (c *consumer) Acknowledge(message *MessageReceiveResponse) error {
	err := c.transporterClient.DeleteMessage(c.queueName, message)
	if err != nil {
		logger.Error(errors.Wrap(err, "failed to delete message"))
		return err
	}

	return nil
}

func (c *consumer) GetWorkerPool() int {
	return c.workerPool
}

type (
	producer struct {
		transporterClient Client
		queueName         string
	}

	ServiceProducer interface {
		Produce(data interface{}) error
	}

	ProducerOption struct {
		TransporterClient Client
		QueueName         string
	}
)

func NewServiceProducer(option *ProducerOption) (ServiceProducer, error) {
	if option == nil {
		return nil, errors.New("option is nil")
	}
	return &producer{
		transporterClient: option.TransporterClient,
		queueName:         option.QueueName,
	}, nil
}

func (p *producer) Produce(data interface{}) error {
	err := p.transporterClient.Publish(&MessagePublishOptions{
		MessageBody: data,
		QueueName:   p.queueName,
	})

	if err != nil {
		logger.Error(errors.Wrap(err, "failed to publish data to the queue"))
		return err
	}

	return nil
}
