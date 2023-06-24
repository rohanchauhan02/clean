package datadog

import (
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/qoala-engineering/qoala-common/log"
	"time"
)

var (
	logger = log.NewQoalaLog()
)

type (
	Client interface {
		AddTag(tag string)
		SendMetric(name string, tags ...string) error
		SendCountMetric(name string, tags ...string) error
		SendDurationMetric(name string, t1, t2 time.Time, tags ...string) error
	}

	datadogClient struct {
		client *statsd.Client
	}

	Config struct {
		Namespace string // project name
		Env       string // dev
		Unit      string // qoalaplus
		Host      string // localhost:8125
	}
)

func NewDataDogClient(c Config) (Client, error) {
	client, err := statsd.New(fmt.Sprintf("%s", c.Host))
	if err != nil {
		return nil, err
	}

	client.Namespace = fmt.Sprintf("%s.", c.Namespace)
	client.Tags = append(client.Tags, fmt.Sprintf("unit:%s", c.Unit))
	client.Tags = append(client.Tags, fmt.Sprintf("environment:%s", c.Env))
	logger.Info("Starting datadog client...")
	return &datadogClient{
		client: client,
	}, nil
}

func (d *datadogClient) AddTag(tag string) {
	d.client.Tags = append(d.client.Tags, tag)
}

func (d *datadogClient) SendMetric(name string, tags ...string) error {
	var t []string
	var err error

	t = append(t, tags...)

	err = d.client.Gauge(name, 1, t, 1)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil

}

func (d *datadogClient) SendCountMetric(name string, tags ...string) error {
	var t []string
	var err error

	t = append(t, tags...)

	err = d.client.Count(name, 1, t, 1)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (d *datadogClient) SendDurationMetric(name string, t1, t2 time.Time, tags ...string) error {
	var t []string
	var err error

	t = append(t, tags...)

	err = d.client.Timing(name, t2.Sub(t1), t, 1)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
