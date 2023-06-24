package datadog

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/rohanchauhan02/common/logs"
)

var (
	logger = logs.NewCommonLog()
)

type Datadog struct {
	client *statsd.Client
}

func NewDatadogClient(c Config) (*Datadog, error) {
	client, err := statsd.New(c.GetHost())
	if err != nil {
		return nil, err
	}
	client.Namespace = fmt.Sprintf("%s.", c.GetNamespace())
	client.Tags = append(client.Tags, fmt.Sprintf("unit:%s", c.GetUnit()))
	client.Tags = append(client.Tags, fmt.Sprintf("environment:%s", c.GetEnv()))
	client.Tags = append(client.Tags, fmt.Sprintf("version:%s", c.GetVersion()))
	logger.Info("Starting datadog client...")
	return &Datadog{
		client: client,
	}, nil
}

func (d Datadog) AddTag(tag string) {
	d.client.Tags = append(d.client.Tags, tag)
}

func (d Datadog) SendMetric(name string, tags ...string) {
	t := []string{}
	t = append(t, tags...)
	d.client.Gauge(name, 1, t, 1)
}

func (d Datadog) SendCountMetric(name string, tags ...string) {
	t := []string{}
	t = append(t, tags...)
	d.client.Count(name, 1, t, 1)
}

func (d Datadog) SendDurationMetric(name string, t1, t2 time.Time, tags ...string) {
	t := []string{}
	t = append(t, tags...)
	d.client.Timing(name, t2.Sub(t1), t, 1)
}
