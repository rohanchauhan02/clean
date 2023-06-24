package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/pkg/errors"
	"github.com/rohanchauhan02/clean/common/datadog"
	"github.com/rohanchauhan02/clean/common/slack"
)

type NotificationType int64

const (
	SLACK NotificationType = iota
	DATADOG

	DatadogPanic        = "status:panic"
	DatadogPanicHandler = "middleware: PanicHandler"
)

func (t NotificationType) GetNotificationType() (string, error) {
	switch t {
	case SLACK:
		return "SLACK", nil
	case DATADOG:
		return "NOTIFICATION", nil
	}

	return "", errors.New("invalid notificaiton type")
}

func (t NotificationType) ToEnumType(s string) NotificationType {
	switch s {
	case "SLACK":
		return SLACK
	case "DATADOG":
		return DATADOG
	}
	// default value for BalanceLevelType
	return SLACK
}

type (
	PanicHandlerOption struct {
		AlertOptions []PanicHandlerAlertOption
		EventName    string
	}

	PanicHandlerAlertOption struct {
		NotificationType NotificationType
		DataDogOption    PanicHandlerDatadogOption
		SlackOption      PanicHandlerSlackOption
	}

	PanicHandlerDatadogOption struct {
		Name   string
		Tag    string
		Client *datadog.Datadog
	}

	PanicHandlerSlackOption struct {
		ChannelID   string
		SlackClient slack.Notification
	}
)

func PanicHandler(option PanicHandlerOption) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			defer func() {
				pRecover := recover()
				if pRecover != nil {
					switch t := pRecover.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("unknown error")
					}

					errMessage := fmt.Sprintf("panic occured in the API with Endpoint- %s \nRoot cause- %s", r.URL.String(), err)
					logger.Error(errMessage)

					publishError(errMessage, option.AlertOptions)

					internalServerResponse(err,
						"PANIC-HANDLER-001",
						errMessage,
						w,
					)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func publishError(errMsg string, options []PanicHandlerAlertOption) {
	for _, option := range options {
		if option.NotificationType == DATADOG {
			datadogOption := option.DataDogOption
			if datadogOption.Name == "" {
				datadogOption.Name = DatadogPanicHandler
			}
			if datadogOption.Tag == "" {
				datadogOption.Tag = DatadogPanic
			}
			tags := []string{datadogOption.Tag}
			datadogOption.Client.SendCountMetric(datadogOption.Name, tags...)
		} else if option.NotificationType == SLACK {
			message := slack.Message{
				Title: "Panic Occurred",
				Body:  fmt.Sprintf("Please check this error: %v", errMsg),
			}

			slackOption := option.SlackOption
			err := slackOption.SlackClient.Error(message)
			if err != nil {
				logger.Errorf(errMsg)
			}
		}
	}
	// else will implement later
}

func GoRoutinePanicHandler(panicHandlerOption PanicHandlerOption) {
	var err error
	pRecover := recover()
	if pRecover != nil {
		debug.PrintStack()
		switch t := pRecover.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		default:
			err = errors.New("unknown error")
		}

		panicMessage := "Panic occured in go routine process"
		if panicHandlerOption.EventName != "" {
			panicMessage = fmt.Sprintf("%v in %v", panicMessage, panicHandlerOption.EventName)
		}
		errMessage := fmt.Sprintf("%v\nRoot cause- %s", panicMessage, err)
		logger.Error(errMessage)

		publishError(errMessage, panicHandlerOption.AlertOptions)
	}
}
