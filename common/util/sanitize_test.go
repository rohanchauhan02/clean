package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeAccents(t *testing.T) {
	inputString := "HALÓ"
	expectedString := "HALO"

	t.Run("sanitize accent string", func(t *testing.T) {
		sanitizedString := SanitizeAccents(inputString)
		assert.Equal(t, expectedString, sanitizedString)
	})
}

func TestSanitizeString(t *testing.T) {
	type TestData struct {
		Input    string
		Expected string
	}

	tests := []TestData{
		{
			Input:    "HALÓ&DUNIA ",
			Expected: "HALO-DUNIA",
		},
		{
			Input:    "FILE__NAME__FINANCE.xlsx",
			Expected: "FILE__NAME__FINANCE.xlsx",
		},
		{
			Input:    "FILE__NAME__FINANCE_BAMBANG_BAMBANG@google.com.xlsx",
			Expected: "FILE__NAME__FINANCE_BAMBANG_BAMBANG@google.com.xlsx",
		},
	}

	t.Run("sanitize string", func(t *testing.T) {
		for _, v := range tests {
			sanitizedString := SanitizeString(v.Input)
			assert.Equal(t, v.Expected, sanitizedString)
		}
	})
}
