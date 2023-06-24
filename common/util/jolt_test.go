package util

import (
	"testing"

	"github.com/rohanchauhan02/clean/common/util/mock"
	"github.com/stretchr/testify/assert"
)

func TestParseJoltMapping(t *testing.T) {
	t.Run("success test parse jolt mapping", func(t *testing.T) {
		parsedByte, err := ParseJoltMapping([]byte(mock.MockSpecJolt), []byte(mock.MockInputJolt))
		if err != nil {
			t.Errorf("failed parse jolt mapping, err: %s", err.Error())
		}

		assert.NotEmpty(t, parsedByte, "success parsed jolt mapping")
	})
}

func TestFailedParseJoltMapping(t *testing.T) {
	t.Run("failed test parse jolt mapping", func(t *testing.T) {
		parsedByte, err := ParseJoltMapping([]byte(mock.MockSpecJolt), []byte(mock.MockInputJoltFailed))
		if err != nil {
			t.Errorf("failed parse jolt mapping, err: %s", err.Error())
		}

		assert.Empty(t, parsedByte, "success test failed parsed jolt mapping")
	})
}
