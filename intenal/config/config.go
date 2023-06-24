package config

import (
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	ImmutableConfig interface {
		GetPort() int
		GetDB() DB
		GetRedis() Redis
		GetAWS() AWS
		GetDatadog() Datadog
		GetSentry() Sentry
		GetApiDoc() ApiDoc
		GetSlack() Slack
		GetSQS() SQS
	}
	Config struct {
		Port    int     `mapstructure:"PORT"`
		DB      DB      `mapstructure:"DB"`
		Redis   Redis   `mapstructure:"Redis"`
		AWS     AWS     `mapstructure:"AWS"`
		Datadog Datadog `mapstructure:"Datadog"`
		Sentry  Sentry  `mapstructure:"Sentry"`
		ApiDoc  ApiDoc  `mapstructure:"ApiDoc"`
		Slack   Slack   `mapstructure:"Slack"`
		SQS     SQS     `mapstructure:"SQS"`
	}
	DB struct {
		Host             string `mapstructure:"HOST"`
		Port             string `mapstructure:"PORT"`
		Name             string `mapstructure:"NAME"`
		User             string `mapstructure:"USER"`
		Password         string `mapstructure:"PASSWORD"`
		MaxIdleConns     int    `mapstructure:"MAX_IDLE_CONNS"`
		MaxOpenConns     int    `mapstructure:"MAX_OPEN_CONNS"`
		MaxLifetimeConns int    `mapstructure:"MAX_LIFETIME_CONNS"`
	}

	Redis struct {
		Host     string `mapstructure:"HOST"`
		Name     string `mapstructure:"NAME"`
		Password string `mapstructure:"PASSWORD"`
	}

	AWS struct {
		AccessKey string `mapstructure:"ACCESS_KEY"`
		SecretKey string `mapstructure:"SECRET_KEY"`
		Region    string `mapstructure:"REGION"`
		Endpoint  string `mapstructure:"ENDPOINT"`
	}

	Datadog struct {
		Namespace   string `mapstructure:"NAMESPACE"`
		ServiceName string `mapstructure:"SERVICE_NAME"`
		ServiceEnv  string `mapstructure:"SERVICE_ENV"`
		Host        string `mapstructure:"HOST"`
		Unit        string `mapstructure:"UNIT"`
	}

	Sentry struct {
		DSN string `mapstructure:"DSN"`
	}

	ApiDoc struct {
		SchemaFilePath string `mapstructure:"SCHEMA_FILE_PATH"`
	}

	Slack struct {
		WebhookURL string `mapstructure:"WEBHOOK_URL"`
	}

	SQS struct {
		AddProduct QueueConfig `mapstructure:"ADD_PRODUCT"`
	}
	QueueConfig struct {
		Name                  string `mapstructure:"NAME"`
		NumberRetrieveMessage int    `mapstructure:"NUMBER_RETRIEVE_MESSAGE"`
		WaitTimeSecond        int    `mapstructure:"WAIT_TIME_SECOND"`
		VisibilityTimeout     int    `mapstructure:"VISIBILITY_TIMEOUT"`
		WorkerPool            int    `mapstructure:"WORKER_POOL"`
	}
)

func (c Config) GetPort() int {
	return c.Port
}

func (c Config) GetDB() DB {
	return c.DB
}

func (c Config) GetRedis() Redis {
	return c.Redis
}

func (c Config) GetAWS() AWS {
	return c.AWS
}

func (c Config) GetDatadog() Datadog {
	return c.Datadog
}

func (c Config) GetSentry() Sentry {
	return c.Sentry
}

func (c Config) GetApiDoc() ApiDoc {
	return c.ApiDoc
}

func (c Config) GetSlack() Slack {
	return c.Slack
}

func (c Config) GetSQS() SQS {
	return c.SQS
}

var (
	confOnce sync.Once
	conf     *Config
)

// NewImmutableConfig is a factory that returns os its config implementation
func NewImmutableConfig() ImmutableConfig {
	confOnce.Do(func() {
		v := viper.New()
		appEnv, exists := os.LookupEnv("APP_ENV")
		configName := "app.config.local"
		if exists {
			switch appEnv {
			case "development":
				configName = "app.config.dev"
			case "uat":
				configName = "app.config.uat"
			case "production":
				configName = "app.config.prod"

			}
		}
		v.SetConfigName("configs/" + configName)
		v.AddConfigPath(".")
		// TO_CHANGE: Change the env prefix based on your service
		v.SetEnvPrefix("GO_CLEAN")
		v.AutomaticEnv()
		if err := v.ReadInConfig(); err != nil {
			panic(err.Error())
		}
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		err := v.Unmarshal(&conf)
		if err != nil {
			panic(err.Error())
		}
	})
	return conf
}
