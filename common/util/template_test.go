package util

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rohanchauhan02/clean/common/util/mock"
)

func TestExecuteTemplateText(t *testing.T) {
	t.Run("test execute template text", func(t *testing.T) {
		payload := &mock.MockPayload{
			FullName: aws.String("fuad"),
		}
		res := ExecuteTemplateText(aws.String(mock.MockTempalteCode), aws.String(mock.MockTempalteCode), payload)
		if res == nil {
			t.Error("failed test execute template text")
		}
	})
}

func TestExecuteTemplateFile(t *testing.T) {
	t.Run("test execute template file", func(t *testing.T) {
		payload := &mock.MockPayload{
			FullName: aws.String("fuad"),
		}
		_, err := ExecuteTemplateFile(mock.MockTemplateFilePath, payload)
		if err != nil {
			t.Error("failed test execute template text")
		}
	})

	t.Run("test execute template file Not Exist", func(t *testing.T) {
		payload := &mock.MockPayload{
			FullName: aws.String("fuad"),
		}
		_, _ = ExecuteTemplateFile(mock.MockTemplateNotExistFilePath, payload)
	})

	t.Run("test execute template file NOK", func(t *testing.T) {
		payload := &mock.MockPayload{
			FullName: aws.String("fuad"),
		}
		_, _ = ExecuteTemplateFile(mock.MockTemplateFailedFilePath, payload)
	})
}
