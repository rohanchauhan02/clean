package mock

const (
	MockAccessKey          = "1234"
	MockSecretKey          = "ABC"
	MockAWSRegion          = "ap-southeast-1"
	MockBucket             = "qoala-mock-testing"
	MockKeyFile            = "/dummy_pdf.pdf"
	MockAWSKeyFileNotExist = "/not_exist/dummy_pdf.pdf"
	MockAWSAccessEmptyKey  = ""
	MockAWSSecretEmptyKey  = ""
	MockAWSRegionEmpty     = ""
	MockAWSEmptyKeyFile    = ""

	MockTemplateFilePath         = "./mock/template.txt"
	MockTemplateFailedFilePath   = "./mock/template_failed.txt"
	MockTemplateNotExistFilePath = "./mock/templateabc.txt"
	MockTempalteCode             = "test-template-code"
	MockTemplate                 = `Hallo {{.FullName}}`

	DocumentType          = "POLICY"
	DocumentFileName      = "policy.pdf"
	UnsanitizedOSSFileName = "private/\tpolicy.pdf"
	SanitizedOSSFileName   = "private/policy.pdf"
	UnsanitizedS3FileName = "/private/\tpolicy.pdf"
	SanitizedS3Filename   = "/private/policy.pdf"
	DocumentSize          = 300
	DocumentUserType      = "USER"
	DocumentMimeType      = "application/pdf"

	ProviderS3       = "S3"
	ProviderGCP      = "GCP"
	ProviderAlicloud = "ALICLOUD"
)

type MockPayload struct {
	FullName *string `json:"fullname" validate:"required"`
}
