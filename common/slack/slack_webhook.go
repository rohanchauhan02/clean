package slack

import (
	"errors"

	"github.com/slack-go/slack"
)

type (
	slackWebhookNotification struct {
		webhookURL string
	}
)

func NewWebhookNotification(webhookURL string) Notification {
	return &slackWebhookNotification{
		webhookURL: webhookURL,
	}
}

func (s slackWebhookNotification) Warning(msg Message) error {
	attachment := createAttachment(msg)
	if attachment == nil {
		return errors.New("failed create attachment")
	}

	attachment.Color = "warning"

	err := slack.PostWebhook(s.webhookURL, &slack.WebhookMessage{
		Attachments: []slack.Attachment{*attachment},
	})

	if err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}

func (s slackWebhookNotification) Info(msg Message) error {
	attachment := createAttachment(msg)
	if attachment == nil {
		return errors.New("failed create attachment")
	}

	attachment.Color = "good"

	err := slack.PostWebhook(s.webhookURL, &slack.WebhookMessage{
		Attachments: []slack.Attachment{*attachment},
	})

	if err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}

func (s slackWebhookNotification) Error(msg Message) error {
	attachment := createAttachment(msg)
	if attachment == nil {
		return errors.New("failed create attachment")
	}

	attachment.Color = "danger"

	err := slack.PostWebhook(s.webhookURL, &slack.WebhookMessage{
		Attachments: []slack.Attachment{*attachment},
	})

	if err != nil {
		logger.Info(err)
		return err
	}

	return nil
}

func (s slackWebhookNotification) CustomChannelNotification(customChannel CustomChannel) error {
	attachment := createAttachment(customChannel.Message)
	if attachment == nil {
		return errors.New("failed create attachment")
	}

	if customChannel.NotificationType == "ERROR" {
		attachment.Color = "danger"
	} else if customChannel.NotificationType == "WARNING" {
		attachment.Color = "warning"
	} else {
		attachment.Color = "good"
	}

	webhookURL := s.webhookURL
	if customChannel.WebhookURL != "" {
		webhookURL = customChannel.WebhookURL
	}
	err := slack.PostWebhook(webhookURL, &slack.WebhookMessage{
		Attachments: []slack.Attachment{*attachment},
	})

	if err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}
