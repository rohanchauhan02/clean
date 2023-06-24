package storage

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rohanchauhan02/clean/common/util"
	"strings"
)

type ClientOptions struct {
	Provider        string `json:"provider"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Region          string `json:"region"`
}

type Client interface {
	CreatePresignedUpload(payload *CreatePresignedUploadRequest) (*CreatePresignedUploadResponse, error)
	CreatePresignedView(payload *CreatePresignedViewRequest) (*CreatePresignedViewResponse, error)
	GetObjectBuffer(payload *GetObjectBufferRequest) ([]byte, error)
	PutObjectBase64(payload *CreateBase64UploadRequest) (*CreateBase64UploadResponse, error)
}

type Options struct {
	Provider        string
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	Region          string
}

//NewClient initializes a new client depending on the provided/vendor type
//The client will be used in the vendor *.go file to access the SDK queue
//methods
func NewClient(options *ClientOptions) (Client, error) {
	provider := strings.ToLower(options.Provider)
	switch provider {
	case "alicloud":
		{
			client, err := oss.New(options.Endpoint, options.AccessKeyID, options.AccessKeySecret)
			if err != nil {
				return nil, err
			}
			oss := &OSS{
				Client: client,
			}
			return oss, nil
		}
	case "aws":
		{
			awsSess, err := util.GetAWSSession(options.AccessKeyID, options.AccessKeySecret, options.Region)
			if err != nil {
				return nil, err
			}
			client := s3.New(awsSess)
			s3 := &S3{
				Client: client,
			}
			return s3, nil
		}
	default:
		{
			return nil, fmt.Errorf("provider:\"%s\" is not supported in storage module", options.Provider)
		}
	}
}
