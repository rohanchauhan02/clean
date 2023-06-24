package storage

import (
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/rohanchauhan02/clean/common/storage/mock"
	"github.com/stretchr/testify/assert"
)

var (
	ossEndpoint       string
	alicloudAccessKey string
	alicloudSecretKey string
	alicloudProvider  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ossEndpoint = os.Getenv("ALICLOUD_OSS_ENDPOINT")
	alicloudAccessKey = os.Getenv("ALICLOUD_ACCESS_KEY")
	alicloudSecretKey = os.Getenv("ALICLOUD_SECRET_KEY")
	alicloudProvider = "alicloud"
}

func TestCreateAlicloudPresignedUpload(t *testing.T) {
	var timeDuration time.Duration

	conf := &ClientOptions{
		Provider:        alicloudProvider,
		Endpoint:        ossEndpoint,
		AccessKeyID:     alicloudAccessKey,
		AccessKeySecret: alicloudSecretKey,
	}
	client, _ := NewClient(conf)

	payload := &CreatePresignedUploadRequest{
		Bucket:   stringpointer(mock.MockBucket),
		Filename: stringpointer(mock.DocumentFileName),
		Mimetype: stringpointer(mock.DocumentMimeType),
		Provider: stringpointer(mock.ProviderAlicloud),
		Size:     int64pointer(mock.DocumentSize),
		Type:     stringpointer(mock.DocumentType),
	}
	//	TODO: (cvs) add mocks and return invalid bucket

	t.Run("test create upload presign url duration OK default duration", func(t *testing.T) {
		resp, err := client.CreatePresignedUpload(payload)
		assert.Nil(t, err, "success create presigned url default duration")
		assert.Contains(t, *resp.URL, "Signature")
		assert.Contains(t, *resp.URL, "Expires")
	})

	t.Run("test create upload presign url with unsanitized filename", func(t *testing.T) {
		unsanitizedPayload := &CreatePresignedUploadRequest{
			Bucket:   stringpointer(mock.MockBucket),
			Filename: stringpointer(mock.UnsanitizedOSSFileName),
			Mimetype: stringpointer(mock.DocumentMimeType),
			Provider: stringpointer(mock.ProviderAlicloud),
			Size:     int64pointer(mock.DocumentSize),
			Type:     stringpointer(mock.DocumentType),
		}
		resp, err := client.CreatePresignedUpload(unsanitizedPayload)
		assert.Nil(t, err, "success create presigned url default duration")
		assert.Equal(t, *resp.Filename, mock.SanitizedOSSFileName)
		assert.Contains(t, *resp.URL, "Signature")
		assert.Contains(t, *resp.URL, "Expires")
	})

	t.Run("test create upload presign url with custom duration", func(t *testing.T) {
		timeDuration = 30 * time.Minute
		payload.Duration = timeDuration
		resp, err := client.CreatePresignedUpload(payload)
		assert.Nil(t, err, "success create presigned url default duration")
		assert.Contains(t, *resp.URL, "Signature")
		assert.Contains(t, *resp.URL, "Expires")
	})

	t.Run("test create with invalid duration", func(t *testing.T) {
		timeDuration = -1 * time.Minute
		payload.Duration = timeDuration
		resp, _ := client.CreatePresignedUpload(payload)
		assert.Nil(t, resp, "error create presigned url with invalid duration")
	})
}

func TestCreateAlicloudSignedView(t *testing.T) {
	var timeDuration time.Duration

	conf := &ClientOptions{
		Provider:        alicloudProvider,
		Endpoint:        ossEndpoint,
		AccessKeyID:     alicloudAccessKey,
		AccessKeySecret: alicloudSecretKey,
	}
	client, _ := NewClient(conf)

	payload := &CreatePresignedViewRequest{
		Bucket: stringpointer(mock.MockBucket),
		Key:    stringpointer(mock.DocumentFileName),
	}

	//	TODO: (cvs) add mocks and return invalid bucket
	t.Run("test create view presigned url duration OK default duration", func(t *testing.T) {
		resp, err := client.CreatePresignedView(payload)
		assert.Nil(t, err, "success making view presigned url default duration")
		assert.Contains(t, *resp.URL, "Signature")
		assert.Contains(t, *resp.URL, "Expires")
	})

	t.Run("test create upload presign url with custom duration", func(t *testing.T) {
		timeDuration = 30 * time.Minute
		payload.Duration = timeDuration
		resp, err := client.CreatePresignedView(payload)
		assert.Nil(t, err, "success making view presigned url default duration")
		assert.Contains(t, *resp.URL, "Signature")
		assert.Contains(t, *resp.URL, "Expires")
	})

	t.Run("test create with invalid duration", func(t *testing.T) {
		timeDuration = -1 * time.Minute
		payload.Duration = timeDuration
		resp, _ := client.CreatePresignedView(payload)
		assert.Nil(t, resp, "error making view presigned url with invalid duration")
	})
}

func TestAlicloudPutObjectBase64(t *testing.T) {

	conf := &ClientOptions{
		Provider:        alicloudProvider,
		AccessKeyID:     alicloudAccessKey,
		AccessKeySecret: alicloudSecretKey,
		Endpoint:        ossEndpoint,
	}
	client, _ := NewClient(conf)

	objectBody := []byte("test_alicloud_put_object_buffer")

	base64Encoded := base64.StdEncoding.EncodeToString(objectBody)
	contentType := http.DetectContentType(objectBody)

	t.Run("base64 upload to alicloud oss happy path", func(t *testing.T) {
		payload := &CreateBase64UploadRequest{
			Filename: stringpointer(mock.DocumentFileName),
			Type:     stringpointer(mock.DocumentType),
			Mimetype: stringpointer(contentType),
			Size:     int64pointer(mock.DocumentSize),
			Bucket:   stringpointer(mock.MockBucket),
			Provider: stringpointer(mock.ProviderAlicloud),
			Base64:   stringpointer(base64Encoded),
		}

		_, err := client.PutObjectBase64(payload)
		assert.Nil(t, err, "failed uploading object to oss")

	})

	t.Run("test upload corrupt object b64", func(t *testing.T) {

		payload := &CreateBase64UploadRequest{
			Filename: aws.String(mock.DocumentFileName),
			Type:     aws.String(mock.DocumentType),
			Mimetype: aws.String(contentType),
			Size:     aws.Int64(mock.DocumentSize),
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Base64:   aws.String(string(objectBody)),
		}

		_, err := client.PutObjectBase64(payload)
		assert.NotNil(t, err, "upload corrupt object b64 failed as expected")
	})

	t.Run("test put object with unsanitized filename", func(t *testing.T) {
		unsanitizedPayload := &CreateBase64UploadRequest{
			Filename: aws.String(mock.UnsanitizedOSSFileName),
			Type:     aws.String(mock.DocumentType),
			Mimetype: aws.String(contentType),
			Size:     aws.Int64(mock.DocumentSize),
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Base64:   stringpointer(base64Encoded),
		}
		resp, err := client.PutObjectBase64(unsanitizedPayload)
		assert.Equal(t, *resp.Filename, mock.SanitizedOSSFileName)
		assert.Nil(t, err, "failed uploading object to oss")
	})
}

func TestAlicloudGetObjectBuffer(t *testing.T) {

	conf := &ClientOptions{
		Provider:        alicloudProvider,
		AccessKeyID:     alicloudAccessKey,
		AccessKeySecret: alicloudSecretKey,
		Endpoint:        ossEndpoint,
	}
	client, _ := NewClient(conf)

	t.Run("test get invalid object from oss bucket", func(t *testing.T) {
		payload := &GetObjectBufferRequest{
			Bucket:   stringpointer(mock.MockBucket),
			Provider: stringpointer(mock.ProviderAlicloud),
			Key:      stringpointer("invalid-object"),
		}

		_, err := client.GetObjectBuffer(payload)
		assert.NotNil(t, err, "failed get invalid object from oss")

	})

	t.Run("test get valid object from oss", func(t *testing.T) {

		payload := &GetObjectBufferRequest{
			Bucket:   stringpointer(mock.MockBucket),
			Provider: stringpointer(mock.ProviderAlicloud),
			Key:      stringpointer("valid-object"),
		}

		_, err := client.GetObjectBuffer(payload)
		assert.NotNil(t, err, "success get valid object from oss")
	})
}
