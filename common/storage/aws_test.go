package storage

import (
	"encoding/base64"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/joho/godotenv"
	"github.com/rohanchauhan02/clean/common/storage/mock"
	"github.com/stretchr/testify/assert"
)

var (
	awsRegion    string
	awsAccessKey string
	awsSecretKey string
	awsProvider  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	awsRegion = mock.MockAWSRegion
	awsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	awsProvider = "aws"
}

func TestCreateS3PresignedUpload(t *testing.T) {
	var timeDuration time.Duration
	conf := &ClientOptions{
		Provider:        awsProvider,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
		Region:          awsRegion,
	}
	client, _ := NewClient(conf)

	payload := &CreatePresignedUploadRequest{
		Bucket:   aws.String(mock.MockBucket),
		Filename: aws.String(mock.DocumentFileName),
		Mimetype: aws.String(mock.DocumentMimeType),
		Provider: aws.String(mock.ProviderS3),
		Size:     aws.Int64(mock.DocumentSize),
		Type:     aws.String(mock.DocumentType),
	}

	t.Run("test create upload presign url duration OK default duration", func(t *testing.T) {
		_, err := client.CreatePresignedUpload(payload)
		assert.Nil(t, err, "success create presigned default duration")
	})

	t.Run("test create presigned upload url", func(t *testing.T) {
		timeDuration = 30 * time.Minute
		payload.Duration = timeDuration

		_, err := client.CreatePresignedUpload(payload)
		assert.Nil(t, err, "success create presigned upload url")
	})

	t.Run("test create upload presign url with unsanitized filename", func(t *testing.T) {
		unsanitizedPayload := &CreatePresignedUploadRequest{
			Bucket:   stringpointer(mock.MockBucket),
			Filename: stringpointer(mock.UnsanitizedS3FileName),
			Mimetype: stringpointer(mock.DocumentMimeType),
			Provider: stringpointer(mock.ProviderAlicloud),
			Size:     int64pointer(mock.DocumentSize),
			Type:     stringpointer(mock.DocumentType),
		}
		resp, err := client.CreatePresignedUpload(unsanitizedPayload)
		assert.Nil(t, err, "success create presigned url default duration")
		assert.Equal(t, *resp.Filename, mock.SanitizedS3Filename)
		assert.Contains(t, *resp.URL, "Signature")
		assert.Contains(t, *resp.URL, "Expires")
	})

	t.Run("test create presigned upload url NOK", func(t *testing.T) {
		timeDuration = -1 * time.Minute
		payload.Duration = timeDuration
		_, err := client.CreatePresignedUpload(payload)
		assert.NotNil(t, err, "create presigned upload url NOK")
	})

}

func TestCreateS3PresignedView(t *testing.T) {
	var timeDuration time.Duration

	conf := &ClientOptions{
		Provider:        awsProvider,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
		Region:          awsRegion,
	}
	client, _ := NewClient(conf)

	payload := &CreatePresignedViewRequest{
		Bucket:   aws.String(mock.MockBucket),
		Provider: aws.String(mock.ProviderS3),
		Key:      aws.String(mock.MockKeyFile),
	}

	t.Run("test create presigned view url default duration", func(t *testing.T) {
		_, err := client.CreatePresignedView(payload)
		assert.Nil(t, err, "success create presigned view url default duration")
	})

	t.Run("test create presigned view url set duration", func(t *testing.T) {
		timeDuration = 1 * time.Minute
		payload.Duration = timeDuration
		_, err := client.CreatePresignedView(payload)
		assert.Nil(t, err, "success create presigned view url set duration")
	})

	t.Run("test create presigned view url set minus duration", func(t *testing.T) {
		timeDuration = -1 * time.Minute
		payload.Duration = timeDuration
		_, err := client.CreatePresignedView(payload)
		assert.NotNil(t, err, "success test presigned view url set minus duration NOK")
	})

}

func TestGetS3ObjectBuffer(t *testing.T) {

	conf := &ClientOptions{
		Provider:        awsProvider,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
		Region:          awsRegion,
	}

	client, _ := NewClient(conf)

	t.Run("test get s3 object buffer ok", func(t *testing.T) {
		payload := &GetObjectBufferRequest{
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Key:      aws.String(mock.MockKeyFile),
		}

		_, err := client.GetObjectBuffer(payload)
		assert.Nil(t, err, "success get s3 object buffer ok")

	})

	t.Run("test get s3 object buffer nok", func(t *testing.T) {
		payload := &GetObjectBufferRequest{
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Key:      aws.String(mock.MockAWSEmptyKeyFile),
		}

		_, err := client.GetObjectBuffer(payload)
		assert.NotNil(t, err, " success s3 object buffer nok")

	})
}

func TestCreateS3ObjectBase64(t *testing.T) {

	conf := &ClientOptions{
		Provider:        awsProvider,
		AccessKeyID:     awsAccessKey,
		AccessKeySecret: awsSecretKey,
		Region:          awsRegion,
	}
	client, _ := NewClient(conf)

	payloadGetObject := &GetObjectBufferRequest{
		Bucket:   aws.String(mock.MockBucket),
		Key:      aws.String(mock.MockKeyFile),
		Provider: aws.String(mock.ProviderS3),
	}

	body, err := client.GetObjectBuffer(payloadGetObject)
	assert.Nil(t, err, "ok get object buffer")

	b64BodyStr := base64.StdEncoding.EncodeToString(body)
	contentType := http.DetectContentType(body)

	t.Run("test ok upload s3 object b64", func(t *testing.T) {

		payload := &CreateBase64UploadRequest{
			Filename: aws.String(mock.DocumentFileName),
			Type:     aws.String(mock.DocumentType),
			Mimetype: aws.String(contentType),
			Size:     aws.Int64(mock.DocumentSize),
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Base64:   aws.String(b64BodyStr),
		}

		_, err = client.PutObjectBase64(payload)
		if err != nil {
			assert.NotNil(t, err, "failed nok upload s3 object b64")
		}

	})

	t.Run("test nok upload corrupt object b64", func(t *testing.T) {

		payload := &CreateBase64UploadRequest{
			Filename: aws.String(mock.DocumentFileName),
			Type:     aws.String(mock.DocumentType),
			Mimetype: aws.String(contentType),
			Size:     aws.Int64(mock.DocumentSize),
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Base64:   aws.String(string(body)),
		}

		_, err = client.PutObjectBase64(payload)
		assert.NotNil(t, err, "success nok upload corrupt object b64")
	})

	t.Run("test nok upload not b64 object", func(t *testing.T) {

		payload := &CreateBase64UploadRequest{
			Filename: aws.String(mock.DocumentFileName),
			Type:     aws.String(mock.DocumentType),
			Mimetype: aws.String(contentType),
			Size:     aws.Int64(mock.DocumentSize),
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Base64:   aws.String(string(body)),
		}

		_, err = client.PutObjectBase64(payload)
		assert.NotNil(t, err, "success nok upload nill object b64")
	})

	t.Run("test unsanitized filename upload object", func(t *testing.T) {
		strBody := base64.StdEncoding.EncodeToString([]byte("unsanitized"))
		unsanitizedPayload := &CreateBase64UploadRequest{
			Filename: aws.String(mock.UnsanitizedS3FileName),
			Type:     aws.String(mock.DocumentType),
			Mimetype: aws.String(contentType),
			Bucket:   aws.String(mock.MockBucket),
			Provider: aws.String(mock.ProviderS3),
			Base64:   aws.String(strBody),
			Size:     aws.Int64(11),
		}

		resp, err := client.PutObjectBase64(unsanitizedPayload)
		assert.Equal(t, *resp.Filename, mock.SanitizedS3Filename)
		assert.Nil(t, err, "failed nok upload s3 object b64")

	})

}
