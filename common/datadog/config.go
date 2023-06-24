package datadog

type Config struct {
	Namespace string // project name
	Env       string // dev
	Unit      string // qoalaplus
	Host      string // localhost:8125
	Version   string // service version or release version
}

func (c Config) GetNamespace() string {
	return c.Namespace
}

func (c Config) GetEnv() string {
	return c.Env
}

func (c Config) GetUnit() string {
	return c.Unit
}

func (c Config) GetHost() string {
	return c.Host
}

func (c Config) GetVersion() string {
	return c.Version
}
