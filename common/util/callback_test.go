package util

import (
	"fmt"
	"testing"

	"github.com/rohanchauhan02/clean/common/util/mock"
)


func TestParseCallbackConfig(t *testing.T) {
	t.Run("test execute callback config success", func(t *testing.T) {
		resp, err := ParseCallbackConfig([]byte(mock.MockSpecJolt), []byte(mock.MockInputJolt), []byte(mock.MockCallbackConfig))
		if err != nil {
			t.Errorf("failed TestParseCallbackConfig, err: %s ", err.Error())
		}
		fmt.Println("heree resp gorequest")
		fmt.Println(resp)
	})
}
