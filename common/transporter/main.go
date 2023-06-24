package transporter

import "github.com/rohanchauhan02/common/logs"

//DefaultConsumerTimeout is the timeout used when calling the
//SDK's API for consuming or receiving a message

// MessageVisibilityTimeout is a timeout on which a message state
// changes from Inactive to Active. Active message means that a message
// can be consumed again. This timeout is used during deleting messages.
const (
	MessageVisibilityTimeout = 5
	WaitTimeSeconds          = 10
)

var (
	logger = logs.NewCommonLog()
)
