package slack

import (
	"encoding/json"
	"strconv"
	"time"

	log "github.com/rohanchauhan02/common/logs"
	"github.com/slack-go/slack"
)

var (
	logger = log.NewCommonLog()
)

type (
	Notification interface {
		Warning(msg Message) error
		Info(msg Message) error
		Error(msg Message) error
		CustomChannelNotification(customChannel CustomChannel) error
	}

	Message struct {
		Title string
		Icon  string
		Body  string
	}

	CustomChannel struct {
		NotificationType string
		WebhookURL       string
		Message          Message
	}
)

func createAttachment(msg Message) *slack.Attachment {
	attachment := &slack.Attachment{
		AuthorName: msg.Title,
		Text:       msg.Body,
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}

	return attachment
}
