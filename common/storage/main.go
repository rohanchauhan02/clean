package storage

import (
	"time"

	log "github.com/rohanchauhan02/common/logs"
)

type CreatePresignedUploadRequest struct {
	Filename *string       `json:"filename"`
	Type     *string       `json:"type"`
	Mimetype *string       `json:"mimetype"`
	Size     *int64        `json:"size"`
	Bucket   *string       `json:"bucket"`
	Provider *string       `json:"provider"`
	Duration time.Duration `json:"duration"` // duration of presigned in seconds
}

type CreatePresignedUploadResponse struct {
	Filename *string `json:"filename"`
	Type     *string `json:"type"`
	Mimetype *string `json:"mimetype"`
	Size     *int64  `json:"size"`
	Bucket   *string `json:"bucket"`
	Provider *string `json:"provider"`
	URL      *string `json:"url"`
	Key      *string `json:"key"`
}

type CreatePresignedViewRequest struct {
	Bucket   *string       `json:"bucket"`
	Provider *string       `json:"provider"`
	Key      *string       `json:"key"`
	Duration time.Duration `json:"duration"` // duration of presigned in seconds
}

type CreatePresignedViewResponse struct {
	Bucket   *string `json:"bucket"`
	Provider *string `json:"provider"`
	Key      *string `json:"key"`
	URL      *string `json:"url"`
}

type GetObjectBufferRequest struct {
	Bucket   *string `json:"bucket"`
	Provider *string `json:"provider"`
	Key      *string `json:"key"`
}

type CreateBase64UploadRequest struct {
	Filename *string `json:"filename"`
	Type     *string `json:"type"`
	Mimetype *string `json:"mimetype"`
	Size     *int64  `json:"size"`
	Bucket   *string `json:"bucket"`
	Provider *string `json:"provider"`
	Base64   *string `json:"base64"`
}

type CreateBase64UploadResponse struct {
	Filename *string `json:"filename"`
	Type     *string `json:"type"`
	Mimetype *string `json:"mimetype"`
	Size     *int64  `json:"size"`
	Bucket   *string `json:"bucket"`
	Provider *string `json:"provider"`
	Status   *bool   `json:"status"`
}

var (
	logger = log.NewCommonLog()
)
