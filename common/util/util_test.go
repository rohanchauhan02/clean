package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rohanchauhan02/clean/common/models"
	"github.com/rohanchauhan02/clean/common/util/mock"
	"gopkg.in/go-playground/validator.v9"
)

func TestRandomString(t *testing.T) {
	randomTypes := []string{
		"UPPERCASE",
		"LOWERCASE",
		"UPPERCASE_ALPHANUMERIC",
		"LOWERCASE_ALPHANUMERIC",
		"DEFAULT",
	}

	for _, randomType := range randomTypes {
		t.Run(fmt.Sprintf("Generate random string %s", randomType), func(t *testing.T) {
			resp := RandomString(3, randomType)
			if resp == "" {
				t.Error("failed return generate random string")
			}
		})
	}

}

func TestRandomStringEngine(t *testing.T) {
	t.Run("test random string engine", func(t *testing.T) {
		var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		resp := randomStringEngine(letter, 3)
		if resp == "" {
			t.Error("failed return random string engine")
		}
	})
}

func TestPaginationCounter(t *testing.T) {
	t.Run("test pagination counter", func(t *testing.T) {
		query := &models.CustomGormPaginationQuery{
			Limit:   10,
			Order:   "asc",
			OrderBy: "created_at",
			Page:    1,
		}
		rows := 100
		meta := PaginationCounter(query, rows)
		if meta == nil {
			t.Error("failed construct meta pagination counter")
		}
	})
}

func TestDebugPrintStruct(t *testing.T) {

	payload := &mock.MockPayload{
		FullName: aws.String("fuad"),
	}

	t.Run("test debug print struct", func(t *testing.T) {
		DebugPrintStruct(&payload)
	})
}

func TestDebugWritePDF(t *testing.T) {
	text := "Hello world"
	textEnc := base64.StdEncoding.EncodeToString([]byte(text))

	t.Run("debug write pdf", func(t *testing.T) {
		err := DebugWritePDF(aws.String(textEnc))
		if err != nil {
			t.Error("failed debug write pdf")
		}
	})

	t.Run("debug write pdf nok", func(t *testing.T) {
		_ = DebugWritePDF(aws.String(text))
	})

}

func TestDebugWriteString(t *testing.T) {
	var text *string

	t.Run("debug write string", func(t *testing.T) {
		text = aws.String("Hello world")
		err := DebugWriteString(text)
		if err != nil {
			t.Error("failed debug write pdf")
		}
	})

}

func TestCustomGormPaginationQuery(t *testing.T) {

	t.Run("test custom pagination query", func(t *testing.T) {

		sqlDB, mock, err := sqlmock.New()
		if err != nil {
			t.Error("failed open sql mock")
		}

		db, err := gorm.Open("mysql", sqlDB)
		if err != nil {
			t.Error("failed open gorm mock")
		}

		mock.ExpectBegin()
		trx := db.Begin()
		_, err = CustomGormPaginationQuery(trx, 10, 1, "created_at", "asc")
		if err != nil {
			t.Error("failed test custom pagination query")
		}
		mock.ExpectCommit()
	})
}

func TestGetAWSSession(t *testing.T) {
	t.Run("test aws session OK", func(t *testing.T) {
		_, err := GetAWSSession(mock.MockAWSAccessKey, mock.MockAWSSecretKey, mock.MockAWSRegion)
		if err != nil {
			t.Error("failed test aws session OK")
		}
	})
	t.Run("test aws session NOK", func(t *testing.T) {
		_, _ = GetAWSSession(mock.MockAWSAccessEmptyKey, mock.MockAWSSecretEmptyKey, mock.MockAWSRegionEmpty)
	})
}

func TestValidate(t *testing.T) {
	cv := CustomValidator{
		Validator: validator.New(),
	}

	t.Run("test validate struct", func(t *testing.T) {

		payload := &mock.MockPayload{
			FullName: aws.String("fuad"),
		}
		err := cv.Validate(payload)
		if err != nil {
			t.Error("failed validate struct")
		}
	})
}

func TestCustomBind(t *testing.T) {
	t.Run("test default custom bind", func(t *testing.T) {
		e := echo.New()
		e.Validator = DefaultValidator()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ca := CustomApplicationContext{
			Context: c,
		}

		payload := &mock.MockPayload{
			FullName: aws.String("fu"),
		}

		err := ca.CustomBind(payload)
		assert.Nil(t, err, "ok test default custom bind")
	})
}

func TestDefaultValidator(t *testing.T) {
	defaultValidator := DefaultValidator()
	payload := &mock.MockPayload{
		FullName: aws.String("fuad"),
	}

	err := defaultValidator.Validate(payload)
	assert.Nil(t, err, "success test default validator")
}

func TestCustomResponse(t *testing.T) {
	t.Run("test custom response", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ca := CustomApplicationContext{
			Context: c,
		}
		err := ca.CustomResponse("success", nil, "success message", http.StatusOK, http.StatusOK, nil)
		if err != nil {
			t.Error("failed test custom response")
		}
	})
}

func TestCustomHTTPErrorHandler(t *testing.T) {
	t.Run("test custom http error handler", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		CustomHTTPErrorHandler(errors.New("ERROR_TEST"), c)

	})

	t.Run("test custom http error handler validation", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cv := CustomValidator{
			Validator: validator.New(),
		}

		payload := &mock.MockPayload{
			FullName: nil,
		}
		err := cv.Validate(payload)

		CustomHTTPErrorHandler(err, c)

	})
}

func TestGenerateQoalaPolicyNumber(t *testing.T) {
	policyNumber := GenerateQoalaPolicyNumber("QOALA")
	assert.NotEmpty(t, policyNumber, "success generate policy number")
}

func TestGetFormattedCurrency(t *testing.T) {
	type TestData struct {
		CurrencyCode   string
		Value          float64
		Precision      int
		ExpectedResult string
	}
	obj := []TestData{
		// IDR
		// Precision -1
		{
			CurrencyCode:   "IDR",
			Value:          1000000.01,
			Precision:      -1,
			ExpectedResult: "Rp.1.000.000",
		},
		// Precision 0
		{
			CurrencyCode:   "IDR",
			Value:          1000000.01,
			Precision:      0,
			ExpectedResult: "Rp.1.000.000",
		},
		// Precision 1
		{
			CurrencyCode:   "IDR",
			Value:          1000000.1,
			Precision:      1,
			ExpectedResult: "Rp.1.000.000,1",
		},
		// Precision 2
		{
			CurrencyCode:   "IDR",
			Value:          1000000.591093,
			Precision:      2,
			ExpectedResult: "Rp.1.000.000,59",
		},
		// Precision 3
		{
			CurrencyCode:   "IDR",
			Value:          1000000.591093,
			Precision:      3,
			ExpectedResult: "Rp.1.000.000,591",
		},
		// Precision 4
		{
			CurrencyCode:   "IDR",
			Value:          100000.0101,
			Precision:      4,
			ExpectedResult: "Rp.100.000,0101",
		},
		// VND
		{
			CurrencyCode:   "VND",
			Value:          100000.102435,
			Precision:      2,
			ExpectedResult: "₫100.000,10",
		},
		// USD
		{
			CurrencyCode:   "USD",
			Value:          100000.102435,
			Precision:      2,
			ExpectedResult: "$100,000.10",
		},
		// JPY
		{
			CurrencyCode:   "JPY",
			Value:          100000.102435,
			Precision:      2,
			ExpectedResult: "¥100,000.10",
		},
		// AUD
		{
			CurrencyCode:   "AUD",
			Value:          100000.102435,
			Precision:      2,
			ExpectedResult: "$100 000.10",
		},
		// PHP
		{
			CurrencyCode:   "PHP",
			Value:          100000.102435,
			Precision:      2,
			ExpectedResult: "₱100,000.10",
		},
	}
	t.Run("test currency formatting", func(t *testing.T) {
		for _, v := range obj {
			result, _, _ := GetFormattedCurrency(v.CurrencyCode, v.Value, v.Precision)
			// Result should not nil
			assert.NotNil(t, result)
			// Should equal to the result from the object
			assert.Equal(t, v.ExpectedResult, result, "should be equal")
		}
	})
}

func TestRounding(t *testing.T) {
	type TestData struct {
		Type           string
		Value          float64
		ExpectedResult float64
	}

	obj := []TestData{
		// Nearest
		{
			Type:           "NEAREST",
			Value:          1.29123,
			ExpectedResult: 1.29,
		},
		{
			Type:           "NEAREST",
			Value:          10890.45010,
			ExpectedResult: 10890.45,
		},
		{
			Type:           "NEAREST",
			Value:          10890.59746,
			ExpectedResult: 10890.60,
		},
		// Round up
		{
			Type:           "UP",
			Value:          1.29123,
			ExpectedResult: 1.30,
		},
		{
			Type:           "UP",
			Value:          10890.45010,
			ExpectedResult: 10890.46,
		},
		{
			Type:           "UP",
			Value:          10890.59746,
			ExpectedResult: 10890.60,
		},
		// Round down
		{
			Type:           "DOWN",
			Value:          1.29123,
			ExpectedResult: 1.29,
		},
		{
			Type:           "DOWN",
			Value:          10890.45010,
			ExpectedResult: 10890.45,
		},
		{
			Type:           "DOWN",
			Value:          10890.59746,
			ExpectedResult: 10890.59,
		},
		// Even
		{
			Type:           "TO_EVEN",
			Value:          1.29123,
			ExpectedResult: 1.00,
		},
		{
			Type:           "TO_EVEN",
			Value:          10890.45010,
			ExpectedResult: 10890.00,
		},
		{
			Type:           "TO_EVEN",
			Value:          10890.59746,
			ExpectedResult: 10891.00,
		},
		// None
		{
			Type:           "NONE",
			Value:          1.29123,
			ExpectedResult: 1.29123,
		},
		{
			Type:           "NONE",
			Value:          10890.45010,
			ExpectedResult: 10890.45010,
		},
		{
			Type:           "NONE",
			Value:          10890.59746,
			ExpectedResult: 10890.59746,
		},
	}

	t.Run("test rounding functionality", func(t *testing.T) {
		for _, v := range obj {
			result := Rounding(v.Type, v.Value)
			// Result should not nil
			assert.NotNil(t, result)
			// Should equal to the result from the object
			assert.Equal(t, v.ExpectedResult, result, "should be equal")
		}
	})
}

func TestGetAge(t *testing.T) {
	type TestData struct {
		dateBirth      string
		today          string
		ExpectedResult int
	}

	obj := []TestData{
		{
			dateBirth:      "1996-01-18",
			today:          "1997-01-18",
			ExpectedResult: 1,
		},
		{
			dateBirth:      "1996-01-18",
			today:          "2000-01-18",
			ExpectedResult: 4,
		},
		{
			dateBirth:      "1996-01-18",
			today:          "2022-01-17",
			ExpectedResult: 25,
		},
		{
			dateBirth:      "1996-01-18",
			today:          "2022-03-16",
			ExpectedResult: 26,
		},
		{
			dateBirth:      "1990-08-15",
			today:          "2020-09-09",
			ExpectedResult: 30,
		},
		{
			dateBirth:      "1990-08-15",
			today:          "2020-08-15",
			ExpectedResult: 30,
		},
		{
			dateBirth:      "1990-08-15",
			today:          "2020-08-14",
			ExpectedResult: 29,
		},
		{
			dateBirth:      "1990-08-15",
			today:          "1991-08-15",
			ExpectedResult: 1,
		},
		{
			dateBirth:      "1990-08-15",
			today:          "1990-08-14",
			ExpectedResult: 0,
		},
	}

	t.Run("test get age functionality", func(t *testing.T) {
		for _, v := range obj {
			// Layout of time (YYYY-MM-DD)
			layout := "2006-01-02"
			dob, _ := time.Parse(layout, v.dateBirth)
			today, _ := time.Parse(layout, v.today)
			result := GetAge(dob, today)
			// Result should not nil
			assert.NotNil(t, result)
			// Should equal to the result from the object
			assert.Equal(t, v.ExpectedResult, result, "should be equal")
		}
	})
}

func TestGetValueJson(t *testing.T) {
	JSONInput := `{"Number":1,"String":"Hello","Array":[1,2,3,4,5,6],"Boolean":true,"Float":1.56,"Object":{"Path":{"To":{"Success":"HARD @ WORK"}}}}`
	type TestData struct {
		Path          string
		Type          string
		ExpectedValue interface{}
	}

	t.Run("test json get value functionality", func(t *testing.T) {
		obj := []TestData{
			{
				Path:          "Number",
				Type:          "NUMBER",
				ExpectedValue: int(1),
			},
			{
				Path:          "Number",
				Type:          "INT64",
				ExpectedValue: int64(1),
			},
			{
				Path:          "Float",
				Type:          "FLOAT",
				ExpectedValue: float32(1.56),
			},
			{
				Path:          "Float",
				Type:          "FLOAT64",
				ExpectedValue: float64(1.56),
			},
			{
				Path:          "Boolean",
				Type:          "BOOLEAN",
				ExpectedValue: true,
			},
			{
				Path:          "String",
				Type:          "STRING",
				ExpectedValue: "Hello",
			},
			{
				Path:          "Object.Path.To.Success",
				Type:          "STRING",
				ExpectedValue: "HARD @ WORK",
			},
			{
				Path: "Object.Path.To",
				Type: "OBJECT",
				ExpectedValue: map[string]interface{}{
					"Success": "HARD @ WORK",
				},
			},
		}

		for _, v := range obj {
			result := GetValueJson(JSONInput, v.Path, v.Type)
			if strings.EqualFold(v.Type, "ARRAY") {
				_, ok := result.([]interface{})
				if !ok {
					assert.Fail(t, fmt.Sprintf("Both result value & type should be equal for value: %v expected value: %v & type: %s expected type: %s", result, v.ExpectedValue, "[]interface{}", v.Type))
				}
			} else if strings.EqualFold(v.Type, "OBJECT") {
				_, ok := result.(map[string]interface{})
				if !ok {
					assert.Fail(t, fmt.Sprintf("Both result value & type should be equal for value: %v expected value: %v & type: %s expected type: %s", result, v.ExpectedValue, "map[string]interface{", v.Type))
				}
			} else {
				assert.Equal(t, v.ExpectedValue, result, fmt.Sprintf("Both result value & type should be equal for value: %v expected value: %v & type: %T expected type: %s", result, v.ExpectedValue, result, v.Type))
			}
		}
	})
}

func TestMysqlRealEscapeString(t *testing.T) {
	test := map[string]string{
		"\\":                                  "\\\\",
		"'":                                   `\'`,
		"\\0":                                 "\\\\0",
		"\n":                                  "\\n",
		"\r":                                  "\\r",
		`"`:                                   `\"`,
		"\x1a":                                "\\Z",
		`<p>123</p><div><img width="1080" />`: `<p>123</p><div><img width=\"1080\" />`,
	}

	for k, v := range test {
		r := MysqlRealEscapeString(k)
		assert.Equal(t, v, r, "should be equal")
	}
}
