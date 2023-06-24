package mock

const (
	MockAWSAccessKey      = "1234"
	MockAWSSecretKey      = "ABC"
	MockAWSRegion         = "ap-southeast-1"
	MockAWSAccessEmptyKey = ""
	MockAWSSecretEmptyKey = ""
	MockAWSRegionEmpty    = ""

	MockTemplateFilePath         = "./mock/template.txt"
	MockTemplateFailedFilePath   = "./mock/template_failed.txt"
	MockTemplateNotExistFilePath = "./mock/templateabc.txt"
	MockTempalteCode             = "test-template-code"
	MockTemplate                 = `Hallo {{.FullName}}`
)

type MockPayload struct {
	FullName *string `json:"fullname" validate:"required"`
}

const MockSpecJolt = `
[
	{
		"operation": "shift",
		"spec": {
			"NoRef": "data.policyNumber",
			"Nilai": "data.premiumAmount",
			"Paket": "data.insurerProduct.package",
			"SourceId": "data.insurerProduct.sourceId",
			"Nama": "data.coveredUsers.fullName",
			"TglLahir": "data.coveredUsers.birthDate",
			"Alamat": "data.coveredUsers.address",
			"Kota": "data.coveredUsers.city",
			"Kode Pos": "data.coveredUsers.postalCode",
			"PhoneNo": "data.coveredUsers.phoneNumber",
			"HandPhone": "data.coveredUsers.mobileNumber",
			"TglMulai": "data.startProtectionAt"
		}
	},
	{
		"operation": "timestamp",
		"spec": {
		"startProtectionAt": {
			"inputFormat": "2019-02-02T01:00:00",
			"outputFormat": "2019-02-02"
		}
		}
	},
	{
		"operation": "default",
		"spec": {
			"Email": "-",
			"Perkerjaan": "-",
			"Kota": "Bandung",
			"AhliWari": [{
				"AhliWarisPersen": "100",
				"AhliWarisNama": "AHLI WARIS YANG SAH",
				"AhliWarisHubungan": "Kerabat"
			}]
		}
	}
]
`

const MockInputJolt = `{
	"data": {
	  "coveredUsers": {
		"fullName": "Muhammad Asnal",
		"birthDate": "",
		"address": "Jln. Raya Merdeka Seberang Tol No. 1",
		"postalCode": 12940
	  },
	  "premiumAmount": 30000,
	  "insurerProduct": {
		"package": "Paket AG",
		"sourceId": "20210309QOALA"
	  },
	  "startProtectionAt": "2019-02-02T00:00:00Z",
	  "endProtectionAt": "2019-03-02T00:00:00Z"
	}
  }`

const MockInputJoltFailed = ``

const MockExpectedOutputJolt = `{"Nama":"Muhammad Asnal","TglLahir":"","Kota":"Bandung","PhoneNo":null,"TglMulai":"2019-02-02T00:00:00Z","NoRef":"PA-TOKPED-005-210222-CQY5F","Nilai":30000,"Paket":"Paket AG","HandPhone":null,"SourceId":"20210309QOALA","Alamat":"Jln. Raya Merdeka Seberang Tol No. 1","Kode Pos":12940,"Email":"-","Perkerjaan":"-","AhliWari":[{"AhliWarisHubungan":"Kerabat","AhliWarisNama":"AHLI WARIS YANG SAH","AhliWarisPersen":"100"}]}`

const MockCallbackConfig = `{
	"method": "POST",
	"base_url": "https://www.example.com",
	"endpoint": "api/broker/notify_claim_info_v2",
	"body_mapping": [
		{
			"operation": "shift",
			"spec": {
				"NoRef": "data.policyNumber",
				"Nilai": "data.premiumAmount",
				"Paket": "data.insurerProduct.package",
				"SourceId": "data.insurerProduct.sourceId",
				"Nama": "data.coveredUsers.fullName",
				"TglLahir": "data.coveredUsers.birthDate",
				"Alamat": "data.coveredUsers.address",
				"Kota": "data.coveredUsers.city",
				"Kode Pos": "data.coveredUsers.postalCode",
				"PhoneNo": "data.coveredUsers.phoneNumber",
				"HandPhone": "data.coveredUsers.mobileNumber",
				"TglMulai": "data.startProtectionAt"
			}
		},
		{
			"operation": "timestamp",
			"spec": {
			"startProtectionAt": {
				"inputFormat": "2019-02-02T01:00:00",
				"outputFormat": "2019-02-02"
			}
			}
		},
		{
			"operation": "default",
			"spec": {
				"Email": "-",
				"Perkerjaan": "-",
				"Kota": "Bandung",
				"AhliWari": [{
					"AhliWarisPersen": "100",
					"AhliWarisNama": "AHLI WARIS YANG SAH",
					"AhliWarisHubungan": "Kerabat"
				}]
			}
		}
	],
	"config_keys": [
	  {
		"name": "shopee_secret_key",
		"key": "SHOPEE_SECRET_KEY"
	  }
	],
	"generated_keys": [
	  {
		"name": "timestamp",
		"type": "timestamp",
		"format": "milliseconds"
	  },
	  {
		"name": "nonce",
		"type": "uuid"
	  },
	  {
		"name": "body_string",
		"type": "body"
	  },
	  {
		"name": "signature",
		"type": "auth",
		"auth_type": "HMAC_SHA256",
		"auth_config": {
		  "message_generation": {
			"format": "%s\n%s\n%s\n%s\n%s\n",
			"params": [
			  "POST",
			  "/api/broker/notify_claim_info_v2",
			  "generated_keys.timestamp",
			  "generated_keys.nonce",
			  "generated_keys.body_string"
			]
		  },
		  "secret": "shopee_secret_key",
		  "encoding": "base64"
		}
	  }
	],
	"headers": {
	  "Content-Type": "application/json",
	  "X-Req-App-Id": "QOALA",
	  "X-Req-Timestamp": "generated_keys.timestamp",
	  "X-Req-Nonce": "generated_keys.nonce",
	  "X-Req-Signature": "generated_keys.signature",
	  "X-Req-Country": "en"
	}
  }`
