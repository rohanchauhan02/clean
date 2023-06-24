package transporter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// GetAWSSession return the AWS session with static credentials or role check
func GetAWSSession(accessKey string, secretKey string, region string) (*session.Session, error) {
	var sess *session.Session
	var err error
	if accessKey != "" && secretKey != "" {
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(region),
		})
	}
	return sess, err
}
