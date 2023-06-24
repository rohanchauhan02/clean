package util

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/derekstavis/go-qs"
	"github.com/fgrosse/goldi"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/leekchan/accounting"
	"github.com/nats-io/nats.go"
	"github.com/rohanchauhan02/clean/common/datadog"
	"github.com/rohanchauhan02/clean/common/models"
	Transporter "github.com/rohanchauhan02/clean/common/transporter"
	"github.com/streadway/amqp"
	"github.com/tidwall/gjson"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
	"gopkg.in/go-playground/validator.v9"
)

type (
	serviceResponse         models.ResponsePattern
	serviceInternalResponse models.ResponseInternalPattern

	// CustomApplicationContext return service custom application context
	CustomApplicationContext struct {
		echo.Context
		Container          *goldi.Container
		RedisSession       *redistrace.Client
		MysqlSession       *gorm.DB
		RabbitSession      *amqp.Connection
		NatsSession        *nats.EncodedConn
		S3Service          *s3.S3
		SqsService         *sqs.SQS
		DynamoService      *dynamodb.DynamoDB
		UserJWT            *models.UserJWT
		DatadogClient      *datadog.Datadog
		TransporterClient  Transporter.Client
		CustomContext      context.Context
		MysqlMasterSession *gorm.DB
	}
)

// CustomBind binding and validate incoming payload request with given struct
func (c *CustomApplicationContext) CustomBind(i interface{}) error {
	if err := bindQueryRequest(c, i); err != nil {
		return err
	}
	err := c.Bind(i)
	if err != nil {
		return err
	}
	err = c.Validate(i)
	if err != nil {
		return err
	}

	return nil
}

func bindQueryRequest(ac *CustomApplicationContext, model interface{}) error {
	queryString := ac.Request().URL.Query().Encode()
	if queryString == "" {
		return nil
	}
	query, err := qs.Unmarshal(queryString)
	if err != nil {
		return err
	}
	jsonString, err := json.Marshal(query)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonString, model)
	if err != nil {
		return err
	}
	return nil
}

// CustomResponse is a method that returns custom object response
func (c *CustomApplicationContext) CustomResponse(status string, data interface{}, message string, code int, httpCode int, meta *models.ResponsePatternMeta) error {

	resp := &serviceResponse{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    code,
		Meta:    meta,
	}

	if meta != nil {
		resp.Meta = meta
	}

	return c.JSON(httpCode, resp)
}

// CustomInternalResponse is a method that returns custom object response
func (c *CustomApplicationContext) CustomInternalResponse(status string, data interface{}, message string, code int, httpCode int, meta *models.ResponseInternalPatternMeta) error {

	resp := &serviceInternalResponse{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    code,
	}

	if meta != nil {
		resp.Meta = meta
	}

	return c.JSON(httpCode, resp)
}

func GetCallerMethod() string {
	var source string
	if pc, _, _, ok := runtime.Caller(2); ok {
		var funcName string
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
			if i := strings.LastIndex(funcName, "."); i != -1 {
				funcName = funcName[i+1:]
			}
		}

		source = path.Base(funcName)
	}
	return source
}

// CustomValidator return  custom validator
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate will validate given input with related struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// DefaultValidator function to give default validation all incoming request
func DefaultValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}

func randomStringEngine(letter []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// RandomString will return random string
func RandomString(n int, kind string) string {
	switch kind {
	case "UPPERCASE":
		var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := randomStringEngine(letter, n)
		return b
	case "LOWERCASE":
		var letter = []rune("abcdefghijklmnopqrstuvwxyz")

		b := randomStringEngine(letter, n)
		return b
	case "UPPERCASE_ALPHANUMERIC":
		var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

		b := randomStringEngine(letter, n)
		return b
	case "LOWERCASE_ALPHANUMERIC":
		var letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
		b := randomStringEngine(letter, n)
		return b
	default:
		var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		b := randomStringEngine(letter, n)
		return b
	}
}

// CustomHTTPErrorHandler will return custom echo http error handler
func CustomHTTPErrorHandler(err error, e echo.Context) {

	report, ok := err.(*echo.HTTPError)
	var msgError string

	if !ok {
		msgError = "[Generic] Internal server error, error [" + err.Error() + "]"
		report = echo.NewHTTPError(http.StatusInternalServerError, msgError)
	}

	var errorCode string
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		msgError = "[Validation] Invalid validation, error [ field: %s is %s ]"
		for _, err := range castedObject {
		b:
			switch err.Tag() {
			case "required":
				msgError = fmt.Sprintf(msgError, err.Field(), "is required")
				report = echo.NewHTTPError(http.StatusBadRequest, msgError)
				errorCode = err.Field() + "_required"
				break b
			}
		}

	}

	e.Logger().Error(msgError)
	qr := &serviceResponse{
		Code:      http.StatusBadRequest,
		Data:      nil,
		Message:   fmt.Sprintf("%+v", report.Message),
		Meta:      nil,
		Status:    "failed",
		ErrorCode: errorCode,
	}

	e.JSON(report.Code, qr)
}

// CustomGormPaginationQuery will return method chaining of gorm fetch pagination
func CustomGormPaginationQuery(trx *gorm.DB, limit int, page int, orderBy string, order string) (*gorm.DB, error) {
	pageOffset := limit * (page - 1)

	if limit != 0 || page != 0 {
		trx = trx.Limit(limit).Offset(pageOffset)
	}
	if orderBy != "" && order != "" {
		trx = trx.Order(fmt.Sprintf("%s %s", orderBy, order))
	}

	return trx, nil
}

// PaginationCounter return response meta counter data
func PaginationCounter(query *models.CustomGormPaginationQuery, rows int) (resp *models.ResponsePatternMeta) {

	meta := &models.ResponsePatternMeta{}

	totalPages := math.Ceil(float64(rows) / float64(query.Limit))
	totalPagesInt := int(totalPages)
	meta.Page = &query.Page
	meta.Limit = &query.Limit
	meta.Count = &rows
	meta.Total = &totalPagesInt
	return meta
}

// DebugPrintStruct print struct to console log
func DebugPrintStruct(input ...interface{}) {
	result, _ := json.Marshal(input)
	fmt.Println(string(result))
}

// DebugWritePDF write PDF from Base64 string to file
func DebugWritePDF(input *string) error {
	dec, err := base64.StdEncoding.DecodeString(*input)
	if err != nil {
		return err
	}

	f, _ := os.Create("debug.pdf")

	defer f.Close()

	_, _ = f.Write(dec)

	_ = f.Sync()

	return nil
}

// DebugWriteString write raw string data to file
func DebugWriteString(input *string) error {
	output := []byte(*input)
	_ = ioutil.WriteFile("debug.txt", output, 0644)

	fmt.Println("[DEBUG PRINT SUCCESS]")
	return nil
}

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

// GenerateQoalaPolicyNumber generate qoala policy number with given prefix
func GenerateQoalaPolicyNumber(prefix string) string {
	currentime := time.Now()

	qoalaPolicySalt := RandomString(8, "UPPERCASE_ALPHANUMERIC")
	qoalaPolicyDate := currentime.Format("20060102")
	qoalaPolicyNumber := prefix + "-" + qoalaPolicyDate + "-" + qoalaPolicySalt

	return qoalaPolicyNumber
}

// GetFormattedCurrency return formatted value with currency symbol i.e
//
//	    American English (USD): formatted: $50,000.50 naked: 50000.5
//	Brazilian Portuguese (BRL): formatted: R$100.000,00 naked: 50000.5
//	              German (EUR): formatted: €50.000,25 naked 50000.25
//	            Japanese (JPY): formatted: ¥50,000 naked 50000
//	               Hindi (INR): formatted: ₹50,000.25 naked 50000.25
//	           Indonesia (IDR): formatted: Rp.50.001 naked 50001
//
// further sample: https://goplay.space/#wseh04IiPxP
// @currencyCode : IDR, USD
// @currencyValue: 50000.00 or 100000.65 or 75000.25 etc
func GetFormattedCurrency(currencyCode string, currencyValue float64, precision int) (formattedValue, formattedNoSymbol string, nakedValue float64) {
	// precision should be positive number
	// if not then reset to 0
	if precision < 0 {
		precision = 0
	}
	lc := accounting.LocaleInfo[currencyCode]
	ac := accounting.Accounting{Symbol: lc.ComSymbol, Precision: precision, Thousand: lc.ThouSep, Decimal: lc.DecSep}
	formatted := ac.FormatMoneyFloat64(currencyValue)
	formattedWithoutSymbol := strings.Replace(formatted, lc.ComSymbol, "", -1)
	accountingFloat64, _ := strconv.ParseFloat(accounting.FormatNumberFloat64(currencyValue, precision, "", "."), 64)

	return formatted, formattedWithoutSymbol, accountingFloat64
}

// BindStruct used to binding source data into target struct
func BindStruct(source interface{}, target interface{}) error {
	sourceBte, err := json.Marshal(&source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(sourceBte, &target)
	if err != nil {
		return err
	}
	return err
}

// Return the rounded value given the value and type (precision 2)
// @roundingType : NEAREST
// @value: 50000.00 or 100000.65 or 75000.25 etc
func Rounding(roundingType string, value float64) float64 {
	var result float64 = 0
	switch strings.ToUpper(roundingType) {
	case "NEAREST":
		result = math.Round(value*100) / 100
	case "UP":
		result = math.Ceil(value*100) / 100
	case "DOWN":
		result = math.Floor(value*100) / 100
	case "TO_EVEN":
		result = math.RoundToEven(value)
	case "NONE":
		result = value
	default:
		result = value
	}
	return result
}

// Return the rounded value given the value and type (precision 2)
// @roundingType : NEAREST
// @value: 50000.00 or 100000.65 or 75000.25 etc
func RoundingNoPrecision(roundingType string, value float64) float64 {
	var result float64 = 0
	switch strings.ToUpper(roundingType) {
	case "NEAREST":
		result = math.Round(value)
	case "UP":
		result = math.Ceil(value)
	case "DOWN":
		result = math.Floor(value)
	case "TO_EVEN":
		result = math.RoundToEven(value)
	case "NONE":
		result = value
	default:
		result = value
	}
	return result
}

// Get age from the two difference date (regardless of leap year)
// Credits: petrus https://forum.golangbridge.org/t/how-to-calculate-the-exact-age-from-given-date-until-today/20530
// See more: https://go.dev/play/p/rdnQBacMH1X
// @birthdate: Time instance time.Parse("1996-01-18") YYYY-MM-DD
// @today: Time instance time.Now() / time.Parse("2022-01-01") YYYY-MM-DD
func GetAge(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

// Get value from JSON encoded string
// @input: string of json encoded data
// @path: path to the property
// @typeData: type data that will be expected to output
// @output: value from json
func GetValueJson(json string, path string, typeData string) (output interface{}) {
	value := gjson.Get(json, path)
	switch strings.ToUpper(typeData) {
	case "NUMBER":
		output = int(value.Int())
	case "INT64":
		output = int64(value.Int())
	case "FLOAT":
		output = float32(value.Float())
	case "FLOAT64":
		output = float64(value.Float())
	case "STRING":
		output = value.String()
	case "BOOLEAN":
		output = value.Bool()
	default:
		output = value.Value()
	}
	return output
}

// Escape the sql injection string
// @value: string String to escape
// @return: safe sql string
// @credits: https://stackoverflow.com/questions/31647406/mysql-real-escape-string-equivalent-for-golang @Shadoweb answer
func MysqlRealEscapeString(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0
	e:
		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break e
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break e
		case '\r':
			escape = 'r'
			break e
		case '\\':
			escape = '\\'
			break e
		case '\'':
			escape = '\''
			break e
		case '"': /* Better safe than sorry */
			escape = '"'
			break e
		case '\032': //十进制26,八进制32,十六进制1a, /* This gives problems on Win32 */
			escape = 'Z'
			break e
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

func ParseStringTimeByLocation(dateTime string, timezone string, dateFormat string) (string, int, error) {
	var utcOffset int
	var err error
	if dateTime == "" {
		return "", utcOffset, errors.New("date time is empty")
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return dateTime, utcOffset, err
	}

	timeParse, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return dateTime, utcOffset, err
	}

	_, utcOffset = timeParse.In(loc).Zone()

	return timeParse.In(loc).Format(dateFormat), utcOffset / 60, nil
}

func ConvertStringDateFilterToTimeByZone(dateFrom string, dateTo string, timezone string) (time.Time, time.Time, error) {
	var utcOffset int
	var from, to time.Time
	var err error

	tm, err := time.LoadLocation(timezone)
	if err != nil {
		return from, to, err
	}
	_, utcOffset = time.Now().In(tm).Zone()

	fromTime, err := time.Parse("2006-01-02", dateFrom)
	if err != nil {
		return from, to, err
	}

	toTime, err := time.Parse("2006-01-02", dateTo)
	if err != nil {
		return from, to, err
	}

	from = fromTime.Add(time.Duration(-utcOffset) * time.Second)
	to = toTime.Add(time.Duration(86399) * time.Second).Add(time.Duration(-utcOffset) * time.Second) // 86399 = 24*3600 - 1 -> to achieve the seconds equivalent for 23:59:59

	return from, to, nil
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
