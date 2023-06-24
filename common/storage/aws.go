package storage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	Client *s3.S3
}

// CreatePresignedUpload used to generate presigned url to upload file based on given input into S3 provider
func (c S3) CreatePresignedUpload(payload *CreatePresignedUploadRequest) (*CreatePresignedUploadResponse, error) {
	if payload.Duration == 0 {
		payload.Duration = DEFAULT_DURATION
	}

	fileKey := sanitizeFileNameForUpload(*payload.Filename)
	mimeType := mime.TypeByExtension(filepath.Ext(*payload.Filename))

	req, _ := c.Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: payload.Bucket,
		Key:    aws.String(fileKey),
	})

	presignedURL, err := req.Presign(payload.Duration)
	if err != nil {
		return nil, err
	}

	resp := &CreatePresignedUploadResponse{
		Filename: aws.String(fileKey),
		Type:     payload.Type,
		Mimetype: aws.String(mimeType),
		Size:     payload.Size,
		Bucket:   payload.Bucket,
		Provider: payload.Provider,
		URL:      aws.String(presignedURL),
		Key:      aws.String(fileKey),
	}
	return resp, nil
}

// CreatePresignedView return s3 object presigned view url with given input
func (c S3) CreatePresignedView(payload *CreatePresignedViewRequest) (*CreatePresignedViewResponse, error) {
	if payload.Duration == 0 {
		payload.Duration = DEFAULT_DURATION
	}

	payload.Key = cleanKey(payload.Key)

	req, _ := c.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: payload.Bucket,
		Key:    payload.Key,
	})

	presignedURL, err := req.Presign(payload.Duration)
	if err != nil {
		return nil, err
	}

	resp := &CreatePresignedViewResponse{
		Bucket:   payload.Bucket,
		Provider: payload.Provider,
		Key:      payload.Key,
		URL:      aws.String(presignedURL),
	}

	return resp, nil
}

// GetObjectBuffer will get object from s3 and return the object as buffer
func (c S3) GetObjectBuffer(payload *GetObjectBufferRequest) ([]byte, error) {
	payload.Key = cleanKey(payload.Key)

	resp, err := c.Client.GetObject(&s3.GetObjectInput{
		Bucket: payload.Bucket,
		Key:    payload.Key,
	})
	if err != nil {
		logger.Error("error during get object for S3 object: ", err)
		return nil, err
	}
	io := resp.Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("error while reading data from reader instance: ", err)
		return nil, err
	}
	defer io.Close()

	return data, nil
}

// PutObjectBase64 upload object to S3 with base64 string data
func (c S3) PutObjectBase64(payload *CreateBase64UploadRequest) (*CreateBase64UploadResponse, error) {

	base64Data, err := base64.StdEncoding.DecodeString(*payload.Base64)
	length := int64(len(base64Data))
	if payload.Size != nil {
		length = *payload.Size
	}
	if err != nil {
		logger.Error(fmt.Printf("Error uploading %s err: %s", *payload.Filename, err.Error()))
		return nil, err
	}

	fileKey := sanitizeFileNameForUpload(*payload.Filename)
	mimeType := mime.TypeByExtension(filepath.Ext(*payload.Filename))

	_, err = c.Client.PutObject(&s3.PutObjectInput{
		ACL:           aws.String("private"),
		Body:          aws.ReadSeekCloser(bytes.NewReader(base64Data)),
		Bucket:        payload.Bucket,
		Key:           aws.String(fileKey),
		ContentType:   aws.String(mimeType),
		ContentLength: &length,
	})

	resp := &CreateBase64UploadResponse{
		Filename: aws.String(fileKey),
		Type:     payload.Type,
		Mimetype: aws.String(mimeType),
		Size:     &length,
		Bucket:   payload.Bucket,
		Provider: payload.Provider,
	}

	if err != nil {
		return nil, err
	}

	resp.Status = aws.Bool(true)
	return resp, nil
}

func cleanKey(key *string) *string {
	if key != nil && strings.HasPrefix(*key, "/") && len(*key) > 1 {
		oriKey := *key
		cleanKey := oriKey[1:]
		return &cleanKey
	}

	return key
}
