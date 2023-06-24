package storage

import "time"

const (
	DEFAULT_DURATION = 1 * time.Hour
	PREFIX_KEY_AWS   = "/private"
	// alicloud PutObject doesn't allow for trailing /,
	// folder names can be referenced with a / such as: private/file.pdf
	PREFIX_KEY_ALICLOUD = "private"
)
