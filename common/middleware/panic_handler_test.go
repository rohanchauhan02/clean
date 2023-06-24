package middleware_test

import (
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/clean/common/middleware"
	"github.com/rohanchauhan02/clean/common/slack"
	"net/http/httptest"
	"testing"
)

type depedency struct {
	WebhookURL string
}

func TestPanicHandler(t *testing.T) {
	testCases := []struct {
		Name      string
		Depedency depedency
	}{
		{
			Name: "it should success!",
			Depedency: depedency{
				WebhookURL: "https://hooks.slack.com/services/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", // # alarms-payment-service-testing
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			slackClient := slack.NewWebhookNotification(tc.Depedency.WebhookURL)
			e := echo.New()
			e.Use(echo.WrapMiddleware(middleware.PanicHandler(middleware.PanicHandlerOption{
				AlertOptions: []middleware.PanicHandlerAlertOption{
					{
						NotificationType: middleware.SLACK,
						SlackOption: middleware.PanicHandlerSlackOption{
							SlackClient: slackClient,
						},
					},
				},
			})))
			e.GET("/", func(context echo.Context) error {
				panic("some event unhandled properly")
			})

			req := httptest.NewRequest(echo.GET, "/", nil)
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)

		})
	}
}

func TestGoRoutinePanicHandler(t *testing.T) {
	testCases := []struct {
		Name      string
		Depedency depedency
	}{
		{
			Name: "it should success!",
			Depedency: depedency{
				WebhookURL: "https://hooks.slack.com/services/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", // # alarms-payment-service-testing
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			slackClient := slack.NewWebhookNotification(tc.Depedency.WebhookURL)

			options := middleware.PanicHandlerOption{
				AlertOptions: []middleware.PanicHandlerAlertOption{
					{
						NotificationType: middleware.SLACK,
						SlackOption: middleware.PanicHandlerSlackOption{
							SlackClient: slackClient,
						},
					},
				},
				EventName: "bulk-update",
			}

			go func() {
				defer middleware.GoRoutinePanicHandler(options)
				panic("some event unhandled properly")
			}()

		})
	}
}
