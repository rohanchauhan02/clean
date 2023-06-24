package storage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSS struct {
	Client *oss.Client
}

// CreatePresignedUpload creates a presigned url for upload using OSS signUrl
func (c OSS) CreatePresignedUpload(payload *CreatePresignedUploadRequest) (*CreatePresignedUploadResponse, error) {
	if payload.Duration == 0 {
		payload.Duration = DEFAULT_DURATION
	}
	bucket, err := c.Client.Bucket(*payload.Bucket)
	if err != nil {
		logger.Error("error while obtaining bucket info ", err)
		return nil, err
	}

	fileName := sanitizeFileNameForUpload(*payload.Filename)
	mimeType := mime.TypeByExtension(filepath.Ext(*payload.Filename))

	options := []oss.Option{
		oss.ContentType(mimeType),
		oss.ContentLength(*payload.Size),
	}

	preSignedURL, err := bucket.SignURL(
		fileName,
		oss.HTTPPut,
		int64(payload.Duration/time.Second),
		options...,
	)

	if err != nil {
		logger.Error("error while generating presigned url ", err)
		return nil, err
	}

	resp := &CreatePresignedUploadResponse{
		Filename: stringpointer(fileName),
		Type:     payload.Type,
		Mimetype: &mimeType,
		Size:     payload.Size,
		Bucket:   payload.Bucket,
		Provider: payload.Provider,
		URL:      &preSignedURL,
		Key:      &fileName,
	}

	return resp, nil

}

// CreatePreSignedView creates a presigned url for viewing an object already present in bucket using OSS signUrl
func (c OSS) CreatePresignedView(payload *CreatePresignedViewRequest) (*CreatePresignedViewResponse, error) {
	if payload.Duration == 0 {
		payload.Duration = DEFAULT_DURATION
	}

	bucket, err := c.Client.Bucket(*payload.Bucket)
	if err != nil {
		logger.Error("error while obtaining bucket info ", err)
		return nil, err
	}

	fileName := *(payload.Key)

	preSignedURL, err := bucket.SignURL(
		fileName,
		oss.HTTPGet,
		int64(payload.Duration/time.Second),
	)

	if err != nil {
		logger.Error("error while generating presigned url ", err)
		return nil, err
	}

	resp := &CreatePresignedViewResponse{
		Bucket:   payload.Bucket,
		Provider: payload.Provider,
		Key:      &fileName,
		URL:      &preSignedURL,
	}

	return resp, nil
}

//GetObjectBuffer returns an object buffer stored on the alicloud OSS bucket
func (c OSS) GetObjectBuffer(payload *GetObjectBufferRequest) ([]byte, error) {
	bucket, err := c.Client.Bucket(*payload.Bucket)
	if err != nil {
		logger.Error("error while obtaining bucket info ", err)
		return nil, err
	}

	io, err := bucket.GetObject(*payload.Key)
	if err != nil {
		logger.Error("error during get object for OSS object: ", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(io)
	if err != nil {
		logger.Error("error while reading data from reader instance: ", err)
		return nil, err
	}
	defer io.Close()
	return data, nil
}

// PutObjectBase64 upload object to alicloud OSS with base64 string data
func (c OSS) PutObjectBase64(payload *CreateBase64UploadRequest) (*CreateBase64UploadResponse, error) {
	base64Data, err := base64.StdEncoding.DecodeString(*payload.Base64)
	length := int64(len(base64Data))
	if payload.Size != nil {
		length = *payload.Size
	}
	if err != nil {
		logger.Error(fmt.Printf("error uploading %s err: %s", *payload.Filename, err.Error()))
		return nil, err
	}

	bucket, err := c.Client.Bucket(*payload.Bucket)
	if err != nil {
		logger.Error("error while obtaining bucket info: ", err)
		return nil, err
	}
	fileKey := sanitizeFileNameForUpload(*payload.Filename)
	mimeType := mime.TypeByExtension(filepath.Ext(*payload.Filename))

	options := []oss.Option{
		oss.ObjectACL(oss.ACLPrivate),
		oss.ContentType(mimeType),
		oss.ContentLength(length),
	}

	err = bucket.PutObject(fileKey, bytes.NewReader(base64Data), options...)
	if err != nil {
		logger.Error("error uploading object: ", err)
		return nil, err
	}

	resp := &CreateBase64UploadResponse{
		Filename: stringpointer(fileKey),
		Type:     payload.Type,
		Mimetype: stringpointer(mimeType),
		Size:     &length,
		Bucket:   payload.Bucket,
		Provider: payload.Provider,
	}
	return resp, nil
}
