package util

import (
	"fmt"
	"testing"

	"github.com/rohanchauhan02/clean/common/schemas"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Config             schemas.UpdateDataConfig
	State              string
	StatusProcess      string
	PreparedStatement  bool
	Expected           []string
	ExpectedTotalQuery int
}

func TestUpdateQueryBuilder(t *testing.T) {
	var test []TestData
	Data := `{"user_channel":"PARTNERSHIP","transaction_number":"TROLOLO-00001-02","policy_number":"KMNL-CRDT-20220519-FN8GA","status":"POLICY_ACTIVE","insurance_code":"SIMAS","category_code":"CRDT","product_code":"KMNL-CRDT","start_protection_at":"2022-05-19T00:00:00Z","end_protection_at":"2022-08-17T00:00:00Z","purchased_at":"2022-05-12T00:00:00Z","currency_code":"IDR","payer_type":"CUSTOMER","calculation":{"gwp":560000,"nett_premium":543200,"admin_fee":0,"hardcopy_admin_fee":0,"marketing_support":0,"after_marketing_support":560000,"total_commission":16800,"total_sum_insured":0,"total_addons":0,"total_tax":0},"addons":[],"policy_holder":{"identity_number":"3175044303960014","identity_type":"KTP","full_name":"PT BAMBANG MAJU BERSAMA","email":"ahmad.fadhel@qoala.id","phone_number":"2127937552","birth_date":"1986-03-03","relationship":"","gender":"","address":"JL. KEMBANG JEPUN 100 SURABAYA 60115","occupation":"","documents":null},"booker":{"identity_number":"","identity_type":"BOOKER","full_name":"PT BAMBANG MAJU BERSAMA","email":"ahmad.fadhel@qoala.id","phone_number":"2127937552","birth_date":"","relationship":"","gender":"","address":"","occupation":"","documents":null},"insureds":[{"type":"LOAN","detail":{"identity_type":"NGANU-001001011","identity_number":"FINDAYA-6-005","loan_amount":80000000,"loan_total_principal":100000000,"loan_total_principal_percentage":80,"tenure_value":90,"tenure_period":"DAYS","interest_rate":19.2,"interest_period":"YEARLY","additional_info":{"fallback_tenure_next_date":91,"premium_rate":0.7}},"documents":[]},{"type":"USER_LENDER","detail":{"identity_number":"3175044303960014","identity_type":"KTP","full_name":"PT KOMUNAL FINANSIAL INDONESIA QQ PT BAMBANG MAJU BERSAMA","email":"ahmad.fadhel@qoala.id","phone_number":"2127937552","birth_date":"1986-03-03","relationship":"","gender":"","address":"JL. KEMBANG JEPUN 100 SURABAYA 60115","occupation":""},"documents":[]},{"type":"USER_BORROWER","detail":{"identity_number":"3578263108880004","identity_type":"KTP","full_name":"PT BAMBANG/BAMBANG","email":"ahmad.fadhel@qoala.id","phone_number":"0859551237612","birth_date":"1979-08-31","relationship":"","gender":"","address":"JL. AHMAD YANI NO. 10 SURABAYA 60234","occupation":"WIRASWASTA"},"documents":[]}],"extra":"","is_standardized":true,"product":{"country":"ID","parent_category":"FINANCE","product_category":"CRDT","product_summary":"","code":"KMNL-CRDT","name":"SIMAS x KOMUNAL Credit Insurance With Reinstatement","category_name":"Credit Insurance with Reinstatement","insurance_name":"Simas Insurtech","insurance_logo":"https://user-service-development.s3.ap-southeast-1.amazonaws.com/documents/public-read/organization/organization-logo/simasnet-1e2eb13fb2f6a6a780502e38ff27e742.jpeg","partner_code":"KOMUNAL","partner_name":"Komunal","partner_logo":"","benefits":[{"code":"CRDT-SIMAS-CI-JY76M","name":"Credit Insurance","benefit_amount":0,"benefit_amount_formatted":"Rp.0,00","logo":"","type":"","description":"Credit insurance is an insurance policy that protects against loss of credit (credit default) and claim submission will be processed directly to the partner after passing the agreed DPD."}],"config":{"insurer_product_config":{"policy":{"activation":{"url":"","fields":null},"expiration":{"unit_duration":"DAYS","unit_duration_amount":0,"fallback_date_type":"FALLBACK_TO_NEXT_DATE","custom_policy_period_field":""},"certificate":{"template_url":"https://policy-service-development.s3.ap-southeast-1.amazonaws.com/templates/finance/template-simas-certificate.html","alternative_template_url":"","wording":"","strategy":"SPLIT","additional_data":null},"numbering":{"type":"RANDOM_ALPHANUMERIC","pad_length":5,"prefix":"product_code"},"callback":{"config":{"method":"POST","base_url":"https://api.simasinsurtech.com/","endpoint":"dataservice_test/asuransi_kredit_json.php","body_mapping":[{"operation":"shift","spec":{"Debitor.[0].Catatan":"","Debitor.[0].DebitorAlamat":"insureds[2].detail.alamat","Debitor.[0].DebitorEmail":"insureds[2].detail.email","Debitor.[0].DebitorHandPhone":"insureds[2].detail.phone_number","Debitor.[0].DebitorIDNo":"insureds[2].detail.identity_type","Debitor.[0].DebitorIDType":"insureds[2].detail.identity_number","Debitor.[0].DebitorKodePos":"","Debitor.[0].DebitorKota":"","Debitor.[0].DebitorName":"insureds[2].detail.full_name","Debitor.[0].DebitorTglLahir":"insureds[2].detail.birth_date","Debitor.[0].InceptionDate":"start_protection_at","Debitor.[0].LamaPinjam":"insureds[0].detail.additional_info.fallback_tenure_next_date","Debitor.[0].Pinjaman":"insureds[0].detail.loan_amount","Debitor.[0].RincianCicilan.[0].Nominal":"insureds[0].detail.loan_amount","Debitor.[0].RincianCicilan.[0].TanggalJatuhTempo":"start_protection_at","KreditorAlamat":"insureds[1].detail.address","KreditorEmail":"insureds[1].detail.email","KreditorHandPhone":"insureds[1].detail.phone_number","KreditorIDNo":"insureds[1].detail.identity_number","KreditorIDType":"insureds[1].detail.identity_type","KreditorKodePos":"","KreditorKota":"","KreditorNama":"insureds[1].detail.full_name","KreditorTglLahir":"insureds[1].detail.birth_date","NoRef":"policy_number"}},{"operation":"default","spec":{"Debitor.[0].IDDebitor":"01","Debitor.[0].Package":"Paket A","SourceID":"20191030KOMUNAL"}},{"operation":"timestamp","spec":{"Debitor.[0].DebitorTglLahir":{"inputFormat":"2006-01-02","outputFormat":"01/02/2006"},"Debitor.[0].InceptionDate":{"inputFormat":"2006-01-02T15:04:05Z07:00","outputFormat":"01/02/2006"},"Debitor.[0].RincianCicilan.[0].TanggalJatuhTempo":{"inputFormat":"2006-01-02T15:04:05Z07:00","outputFormat":"01/02/2006"},"KreditorTglLahir":{"inputFormat":"2006-01-02","outputFormat":"01/02/2006"}}},{"operation":"concat","spec":{"delim":":","sources":[{"path":"KreditorIDType"},{"path":"KreditorIDNo"}],"targetPath":"KreditorIDNo"}},{"operation":"concat","spec":{"delim":":","sources":[{"path":"Debitor.[0].DebitorIDNo"},{"path":"Debitor.[0].DebitorIDType"}],"targetPath":"Debitor.[0].DebitorIDNo"}},{"operation":"delete","spec":{"paths":["Debitor.[0].DebitorIDType","KreditorIDType"]}}],"generated_keys":[{"name":"token","type":"auth","auth_type":"REST","auth_config":{}}],"headers":{"Authorization":"generated_keys.token.response_mapping.bearer_token","Content-Type":"application/json"},"success_response_body":{"ErrorCode":[0]},"update_data_config":{"trigger_status":{"POLICY_ACTIVE":"POLICY_ACTIVE"},"mappings":[{"name":"POLICY_ACTIVE","table":"policies","status":"success","conditions":[{"key":"policy_number = ?","value":"policy_number"}],"column_mappings":[{"key":"insurance_number","type":"string","operation":"gjson","value":"0.certificate"},{"key":"status","type":"string","operation":"static","value":"POLICY_ACTIVE"}]},{"name":"POLICY_ACTIVE","table":"policies","status":"failed","conditions":[{"key":"policy_number = ?","value":"policy_number"}],"column_mappings":[{"key":"status","type":"string","operation":"static","value":"POLICY_CANCELLED"}]},{"name":"POLICY_ACTIVE","table":"callback_logs","status":"json_operation","conditions":[{"key":"policy_number = ?","value":"policy_number"}],"column_mappings":[{"key":"raw_response","type":"string","operation":"json","value":"0.status"}]}]}},"statuses":{"POLICY_ACTIVE":"policy_active"},"status_journey":{"POLICY_ACTIVE":[]},"start_status":["POLICY_ACTIVE"]},"product_validation_config":{"validation":[{"type":"RANGE","parameter":"PREMIUM_UNIT_DURATION","min":1,"max":365,"value_duration":"DAYS"}],"additional_data":null},"start_protection":{"unit_duration":"DAYS","fallback_date_type":"FALLBACK_TO_PREVIOUS_DATE"}},"claim":{"max_limitation":1,"journey_end":"QOALA_CLAIM_APPROVE","claim_period_validation_skip":true,"claim_deductible":{"unit_fee":"PERCENTAGE","unit_fee_amount":0},"claim_expiration":{"unit_duration":"DAYS","unit_duration_amount":0},"callback":{"config":{"method":"POST","base_url":"https://api.simasinsurtech.com/","endpoint":"dataservice_test/klaim_asuransi_kredit_json.php","body_mapping":[{"operation":"shift","spec":{"LamaTelatBayar":"additional_info.overdue_payment_duration","NamaBank":"bank_name","NamaPemilikRekening":"bank_account_name","NilaiKlaim":"net_claim_amount","NoRef":"policy_details.transaction_id","NoRekeningBank":"bank_account_number","PolicyInsuranceNo":"policy_details.insurance_number"}},{"operation":"default","spec":{"SourceID":"20191030KOMUNAL"}},{"operation":"default","spec":{"IDDebitor":"01"}},{"operation":"default","spec":{"Catatan":""}}],"generated_keys":[{"name":"token","type":"auth","auth_type":"REST","auth_config":{}}],"headers":{"Authorization":"generated_keys.token.response_mapping.bearer_token","Content-Type":"application/json"},"success_response_body":{"ErrorCode":[0]}},"statuses":{"CLAIM_INITIATED":"CLAIM_INITIATED"},"status_journey":{"CLAIM_INITIATED":[]},"start_status":["CLAIM_INITIATED"]},"claim_product_validation_flag":true,"product_validation_config":{"statuses":{"CLAIM_INITIATE":true},"validation":[{"type":"RANGE","parameter":"OVERDUE_LIMIT","min":90,"max":90,"value_duration":"DAYS"}],"additional_data":null},"claim_amount_calculation_config":{"need_excess_fee":false,"balance_level":"POLICY","remaining_balance_calculation_statuses":{}}}},"partner_product_config":{"notifications":{"CLAIM_INITIATE":[],"INSURANCE_CLAIM_APPROVE":[],"INSURANCE_CLAIM_PAID":[],"INSURANCE_CLAIM_REJECT":[],"POLICY_ISSUED":[],"QOALA_CLAIM_APPROVE":[],"QOALA_CLAIM_REJECT":[]},"commissions":[{"type":"PERCENTAGE","value":0,"recipient":"PARTNER"},{"type":"PERCENTAGE","value":3,"recipient":"QOALA"}],"payment":{"payment_handler":"PAID","term_of_payment":0},"policy":{"callback":{"config":{}}},"claim":{"callback":{"config":{}}},"stoploss":{"record":false},"additional_config":{"partner_name":"PT KOMUNAL FINANSIAL INDONESIA","policy_holder":{"format":"PARTNER_NAME QQ LENDER_NAME","type":"LENDER_BENEFICIARY_NAME"}}}}},"documents":[{"type":"POLICY_CERTIFICATE","filename":"KMNL-CRDT-20220519-FN8GA.pdf","alias":"policy_certificate.pdf","url":"https://api-staging.qoala.app/api/v2/policies/documents/3d80403c-149a-49a0-a5e5-c8356bf8619e/files/KMNL-CRDT-20220519-FN8GA.pdf"}],"created_at":"2022-05-19T08:33:29Z","payment_at":"2022-05-19T08:33:29Z","commissions":null}`
	Response := `[{"status":"200","auth":"d343b3afb5544cc1c0eef7ee3cc8d830","certificate":"SP2112152000002-001377","nik":"3175064805941002","contract":"A15857304470085143","rate":1,"premi":36000,"urlcert":"https://intraasia.id/apidata.dev/kredit/certificateqoala?certificate=SP2112152000002-001377","result":"Success","MaliciousString":"HEY'DROP TABLE POLICIES;'"}]`

	config := schemas.UpdateDataConfig{
		TriggerStatus: map[string]string{
			"POLICY_INITIATED": "UUID1",
			"CLAIM_INITIATED":  "UUID2",
			"JSON_OPERATIONS":  "UUID3",
			"SQL_INJECTION":    "UUID4",
		},
		Mappings: []*schemas.UpdateDataMapping{
			// UUID1 SUCCESS
			{
				Name:   "UUID1",
				Table:  "policies",
				Status: "success",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
					{
						Key:   "transaction_number = ?",
						Value: "transaction_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "insurance_number",
						Operation: "gjson",
						Type:      "string",
						Value:     "0.certificate",
					},
					{
						Key:       "updated_at",
						Operation: "static",
						Type:      "function",
						Value:     "NOW()",
					},
				},
			},
			{
				Name:   "UUID1",
				Table:  "callback_logs",
				Status: "success",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "raw_response",
						Operation: "json",
						Type:      "int",
						Value:     "0.status",
					},
				},
			},
			// UUID1 FAILED
			{
				Name:   "UUID1",
				Table:  "policies",
				Status: "failed",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
					{
						Key:   "transaction_number = ?",
						Value: "transaction_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "status",
						Operation: "static",
						Type:      "string",
						Value:     "INSURANCE_REJECTED",
					},
					{
						Key:       "updated_at",
						Operation: "static",
						Type:      "function",
						Value:     "NOW()",
					},
				},
			},

			// UUID2 SUCCESS
			{
				Name:   "UUID2",
				Table:  "claims",
				Status: "success",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "last_status",
						Operation: "static",
						Type:      "string",
						Value:     "QOALA_CLAIM_APPROVE",
					},
					{
						Key:       "updated_at",
						Operation: "static",
						Type:      "function",
						Value:     "NOW()",
					},
				},
			},
			// UUID2 FAILED
			{
				Name:   "UUID2",
				Table:  "claims",
				Status: "failed",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "last_status",
						Operation: "static",
						Type:      "string",
						Value:     "QOALA_CLAIM_REJECT",
					},
					{
						Key:       "updated_at",
						Operation: "static",
						Type:      "function",
						Value:     "NOW()",
					},
				},
			},
			// JSON OPERATIONS
			{
				Name:   "UUID3",
				Table:  "callback_logs",
				Status: "success",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "raw_response",
						Operation: "json",
						Type:      "int",
						Value:     "0.status",
					},
				},
			},
			// SQL_INJECTION
			{
				Name:   "UUID4",
				Table:  "callback_logs",
				Status: "success",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "raw_response",
						Operation: "static",
						Type:      "string",
						Value:     `Hello"DROP table policies;"`,
					},
				},
			},
			{
				Name:   "UUID4",
				Table:  "callback_logs",
				Status: "failed",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "raw_response",
						Operation: "json",
						Type:      "string",
						Value:     "0.MaliciousString",
					},
				},
			},
		},
	}

	test = append(test, []TestData{
		// Raw Queries
		{
			ExpectedTotalQuery: 2,
			PreparedStatement:  false,
			State:              "POLICY_INITIATED",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				"UPDATE policies SET insurance_number = 'SP2112152000002-001377', updated_at = NOW() WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA' AND transaction_number = 'TROLOLO-00001-02'",
				"UPDATE callback_logs SET raw_response = JSON_SET(raw_response, '$.0.status', 200) WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "POLICY_INITIATED",
			StatusProcess:      "failed",
			Config:             config,
			Expected: []string{
				"UPDATE policies SET status = 'INSURANCE_REJECTED', updated_at = NOW() WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA' AND transaction_number = 'TROLOLO-00001-02'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "CLAIM_INITIATED",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				"UPDATE claims SET last_status = 'QOALA_CLAIM_APPROVE', updated_at = NOW() WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "CLAIM_INITIATED",
			StatusProcess:      "failed",
			Config:             config,
			Expected: []string{
				"UPDATE claims SET last_status = 'QOALA_CLAIM_REJECT', updated_at = NOW() WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "JSON_OPERATIONS",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				"UPDATE callback_logs SET raw_response = JSON_SET(raw_response, '$.0.status', 200) WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "SQL_INJECTION",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				"UPDATE callback_logs SET raw_response = 'Hello\\\"DROP table policies;\\\"' WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "SQL_INJECTION",
			StatusProcess:      "failed",
			Config:             config,
			Expected: []string{
				`UPDATE callback_logs SET raw_response = JSON_SET(raw_response, '$.0.MaliciousString', 'HEY\'DROP TABLE POLICIES;\'') WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'`,
			},
		},

		// Prepared statements
		{
			ExpectedTotalQuery: 2,
			PreparedStatement:  true,
			State:              "POLICY_INITIATED",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdatepolicies FROM 'UPDATE policies SET insurance_number = ?, updated_at = ? WHERE policy_number = ? AND transaction_number = ?';SET @Values0 = "0.certificate";SET @Values1 = "NOW()";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";SET @Cond1 = "TROLOLO-00001-02";EXECUTE stmntupdatepolicies USING @Values0, @Values1, @Cond0, @Cond1;DEALLOCATE PREPARE stmntupdatepolicies;`,
				`PREPARE stmntupdatecallbacklogs FROM 'UPDATE callback_logs SET raw_response = ? WHERE policy_number = ?';SET @Values0 = "JSON_SET(raw_response, '$.0.status', 200)";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";EXECUTE stmntupdatecallbacklogs USING @Values0, @Cond0;DEALLOCATE PREPARE stmntupdatecallbacklogs;`,
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  true,
			State:              "POLICY_INITIATED",
			StatusProcess:      "failed",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdatepolicies FROM 'UPDATE policies SET status = ?, updated_at = ? WHERE policy_number = ? AND transaction_number = ?';SET @Values0 = "INSURANCE_REJECTED";SET @Values1 = "NOW()";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";SET @Cond1 = "TROLOLO-00001-02";EXECUTE stmntupdatepolicies USING @Values0, @Values1, @Cond0, @Cond1;DEALLOCATE PREPARE stmntupdatepolicies;`,
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  true,
			State:              "CLAIM_INITIATED",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdateclaims FROM 'UPDATE claims SET last_status = ?, updated_at = ? WHERE policy_number = ?';SET @Values0 = "QOALA_CLAIM_APPROVE";SET @Values1 = "NOW()";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";EXECUTE stmntupdateclaims USING @Values0, @Values1, @Cond0;DEALLOCATE PREPARE stmntupdateclaims;`,
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  true,
			State:              "CLAIM_INITIATED",
			StatusProcess:      "failed",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdateclaims FROM 'UPDATE claims SET last_status = ?, updated_at = ? WHERE policy_number = ?';SET @Values0 = "QOALA_CLAIM_REJECT";SET @Values1 = "NOW()";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";EXECUTE stmntupdateclaims USING @Values0, @Values1, @Cond0;DEALLOCATE PREPARE stmntupdateclaims;`,
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  true,
			State:              "JSON_OPERATIONS",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdatecallbacklogs FROM 'UPDATE callback_logs SET raw_response = ? WHERE policy_number = ?';SET @Values0 = "JSON_SET(raw_response, '$.0.status', 200)";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";EXECUTE stmntupdatecallbacklogs USING @Values0, @Cond0;DEALLOCATE PREPARE stmntupdatecallbacklogs;`,
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  true,
			State:              "SQL_INJECTION",
			StatusProcess:      "success",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdatecallbacklogs FROM 'UPDATE callback_logs SET raw_response = ? WHERE policy_number = ?';SET @Values0 = "Hello\"DROP table policies;\"";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";EXECUTE stmntupdatecallbacklogs USING @Values0, @Cond0;DEALLOCATE PREPARE stmntupdatecallbacklogs;`,
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  true,
			State:              "SQL_INJECTION",
			StatusProcess:      "failed",
			Config:             config,
			Expected: []string{
				`PREPARE stmntupdatecallbacklogs FROM 'UPDATE callback_logs SET raw_response = ? WHERE policy_number = ?';SET @Values0 = "JSON_SET(raw_response, '$.0.MaliciousString', HEY\'DROP TABLE POLICIES;\')";SET @Cond0 = "KMNL-CRDT-20220519-FN8GA";EXECUTE stmntupdatecallbacklogs USING @Values0, @Cond0;DEALLOCATE PREPARE stmntupdatecallbacklogs;`,
			},
		},
	}...)

	t.Run("Check UpdateQueryBuilder capability", func(t *testing.T) {
		for _, v := range test {
			queries, err := UpdateQueryBuilder(v.Config, Data, Response, v.State, v.StatusProcess, v.PreparedStatement)
			if err != nil {
				assert.Fail(t, fmt.Sprintf("error occured. err: %v", err.Error()))
			}
			// Add check total query
			if v.ExpectedTotalQuery != 0 {
				assert.Equal(t, v.ExpectedTotalQuery, len(queries), "Should equal")
			}
			for k, query := range queries {
				assert.Equal(t, v.Expected[k], query, "Query should equal each other")
			}
		}
	})
}

func TestUpdateQueryBuilderWithRegex(t *testing.T) {
	var test []TestData
	Data := `{"user_channel":"PARTNERSHIP","transaction_number":"TROLOLO-00001-02","policy_number":"KMNL-CRDT-20220519-FN8GA","status":"POLICY_ACTIVE","insurance_code":"SIMAS","category_code":"CRDT","product_code":"KMNL-CRDT","start_protection_at":"2022-05-19T00:00:00Z","end_protection_at":"2022-08-17T00:00:00Z","purchased_at":"2022-05-12T00:00:00Z","currency_code":"IDR","payer_type":"CUSTOMER","calculation":{"gwp":560000,"nett_premium":543200,"admin_fee":0,"hardcopy_admin_fee":0,"marketing_support":0,"after_marketing_support":560000,"total_commission":16800,"total_sum_insured":0,"total_addons":0,"total_tax":0},"addons":[],"policy_holder":{"identity_number":"3175044303960014","identity_type":"KTP","full_name":"PT BAMBANG MAJU BERSAMA","email":"ahmad.fadhel@qoala.id","phone_number":"2127937552","birth_date":"1986-03-03","relationship":"","gender":"","address":"JL. KEMBANG JEPUN 100 SURABAYA 60115","occupation":"","documents":null},"booker":{"identity_number":"","identity_type":"BOOKER","full_name":"PT BAMBANG MAJU BERSAMA","email":"ahmad.fadhel@qoala.id","phone_number":"2127937552","birth_date":"","relationship":"","gender":"","address":"","occupation":"","documents":null},"insureds":[{"type":"LOAN","detail":{"identity_type":"NGANU-001001011","identity_number":"FINDAYA-6-005","loan_amount":80000000,"loan_total_principal":100000000,"loan_total_principal_percentage":80,"tenure_value":90,"tenure_period":"DAYS","interest_rate":19.2,"interest_period":"YEARLY","additional_info":{"fallback_tenure_next_date":91,"premium_rate":0.7}},"documents":[]},{"type":"USER_LENDER","detail":{"identity_number":"3175044303960014","identity_type":"KTP","full_name":"PT KOMUNAL FINANSIAL INDONESIA QQ PT BAMBANG MAJU BERSAMA","email":"ahmad.fadhel@qoala.id","phone_number":"2127937552","birth_date":"1986-03-03","relationship":"","gender":"","address":"JL. KEMBANG JEPUN 100 SURABAYA 60115","occupation":""},"documents":[]},{"type":"USER_BORROWER","detail":{"identity_number":"3578263108880004","identity_type":"KTP","full_name":"PT BAMBANG/BAMBANG","email":"ahmad.fadhel@qoala.id","phone_number":"0859551237612","birth_date":"1979-08-31","relationship":"","gender":"","address":"JL. AHMAD YANI NO. 10 SURABAYA 60234","occupation":"WIRASWASTA"},"documents":[]}],"extra":"","is_standardized":true,"product":{"country":"ID","parent_category":"FINANCE","product_category":"CRDT","product_summary":"","code":"KMNL-CRDT","name":"SIMAS x KOMUNAL Credit Insurance With Reinstatement","category_name":"Credit Insurance with Reinstatement","insurance_name":"Simas Insurtech","insurance_logo":"https://user-service-development.s3.ap-southeast-1.amazonaws.com/documents/public-read/organization/organization-logo/simasnet-1e2eb13fb2f6a6a780502e38ff27e742.jpeg","partner_code":"KOMUNAL","partner_name":"Komunal","partner_logo":"","benefits":[{"code":"CRDT-SIMAS-CI-JY76M","name":"Credit Insurance","benefit_amount":0,"benefit_amount_formatted":"Rp.0,00","logo":"","type":"","description":"Credit insurance is an insurance policy that protects against loss of credit (credit default) and claim submission will be processed directly to the partner after passing the agreed DPD."}],"config":{"insurer_product_config":{"policy":{"activation":{"url":"","fields":null},"expiration":{"unit_duration":"DAYS","unit_duration_amount":0,"fallback_date_type":"FALLBACK_TO_NEXT_DATE","custom_policy_period_field":""},"certificate":{"template_url":"https://policy-service-development.s3.ap-southeast-1.amazonaws.com/templates/finance/template-simas-certificate.html","alternative_template_url":"","wording":"","strategy":"SPLIT","additional_data":null},"numbering":{"type":"RANDOM_ALPHANUMERIC","pad_length":5,"prefix":"product_code"},"callback":{"config":{"method":"POST","base_url":"https://api.simasinsurtech.com/","endpoint":"dataservice_test/asuransi_kredit_json.php","body_mapping":[{"operation":"shift","spec":{"Debitor.[0].Catatan":"","Debitor.[0].DebitorAlamat":"insureds[2].detail.alamat","Debitor.[0].DebitorEmail":"insureds[2].detail.email","Debitor.[0].DebitorHandPhone":"insureds[2].detail.phone_number","Debitor.[0].DebitorIDNo":"insureds[2].detail.identity_type","Debitor.[0].DebitorIDType":"insureds[2].detail.identity_number","Debitor.[0].DebitorKodePos":"","Debitor.[0].DebitorKota":"","Debitor.[0].DebitorName":"insureds[2].detail.full_name","Debitor.[0].DebitorTglLahir":"insureds[2].detail.birth_date","Debitor.[0].InceptionDate":"start_protection_at","Debitor.[0].LamaPinjam":"insureds[0].detail.additional_info.fallback_tenure_next_date","Debitor.[0].Pinjaman":"insureds[0].detail.loan_amount","Debitor.[0].RincianCicilan.[0].Nominal":"insureds[0].detail.loan_amount","Debitor.[0].RincianCicilan.[0].TanggalJatuhTempo":"start_protection_at","KreditorAlamat":"insureds[1].detail.address","KreditorEmail":"insureds[1].detail.email","KreditorHandPhone":"insureds[1].detail.phone_number","KreditorIDNo":"insureds[1].detail.identity_number","KreditorIDType":"insureds[1].detail.identity_type","KreditorKodePos":"","KreditorKota":"","KreditorNama":"insureds[1].detail.full_name","KreditorTglLahir":"insureds[1].detail.birth_date","NoRef":"policy_number"}},{"operation":"default","spec":{"Debitor.[0].IDDebitor":"01","Debitor.[0].Package":"Paket A","SourceID":"20191030KOMUNAL"}},{"operation":"timestamp","spec":{"Debitor.[0].DebitorTglLahir":{"inputFormat":"2006-01-02","outputFormat":"01/02/2006"},"Debitor.[0].InceptionDate":{"inputFormat":"2006-01-02T15:04:05Z07:00","outputFormat":"01/02/2006"},"Debitor.[0].RincianCicilan.[0].TanggalJatuhTempo":{"inputFormat":"2006-01-02T15:04:05Z07:00","outputFormat":"01/02/2006"},"KreditorTglLahir":{"inputFormat":"2006-01-02","outputFormat":"01/02/2006"}}},{"operation":"concat","spec":{"delim":":","sources":[{"path":"KreditorIDType"},{"path":"KreditorIDNo"}],"targetPath":"KreditorIDNo"}},{"operation":"concat","spec":{"delim":":","sources":[{"path":"Debitor.[0].DebitorIDNo"},{"path":"Debitor.[0].DebitorIDType"}],"targetPath":"Debitor.[0].DebitorIDNo"}},{"operation":"delete","spec":{"paths":["Debitor.[0].DebitorIDType","KreditorIDType"]}}],"generated_keys":[{"name":"token","type":"auth","auth_type":"REST","auth_config":{}}],"headers":{"Authorization":"generated_keys.token.response_mapping.bearer_token","Content-Type":"application/json"},"success_response_body":{"ErrorCode":[0]},"update_data_config":{"trigger_status":{"POLICY_ACTIVE":"POLICY_ACTIVE"},"mappings":[{"name":"POLICY_ACTIVE","table":"policies","status":"success","conditions":[{"key":"policy_number = ?","value":"policy_number"}],"column_mappings":[{"key":"insurance_number","type":"string","operation":"gjson","value":"0.certificate"},{"key":"status","type":"string","operation":"static","value":"POLICY_ACTIVE"}]},{"name":"POLICY_ACTIVE","table":"policies","status":"failed","conditions":[{"key":"policy_number = ?","value":"policy_number"}],"column_mappings":[{"key":"status","type":"string","operation":"static","value":"POLICY_CANCELLED"}]},{"name":"POLICY_ACTIVE","table":"callback_logs","status":"json_operation","conditions":[{"key":"policy_number = ?","value":"policy_number"}],"column_mappings":[{"key":"raw_response","type":"string","operation":"json","value":"0.status"}]}]}},"statuses":{"POLICY_ACTIVE":"policy_active"},"status_journey":{"POLICY_ACTIVE":[]},"start_status":["POLICY_ACTIVE"]},"product_validation_config":{"validation":[{"type":"RANGE","parameter":"PREMIUM_UNIT_DURATION","min":1,"max":365,"value_duration":"DAYS"}],"additional_data":null},"start_protection":{"unit_duration":"DAYS","fallback_date_type":"FALLBACK_TO_PREVIOUS_DATE"}},"claim":{"max_limitation":1,"journey_end":"QOALA_CLAIM_APPROVE","claim_period_validation_skip":true,"claim_deductible":{"unit_fee":"PERCENTAGE","unit_fee_amount":0},"claim_expiration":{"unit_duration":"DAYS","unit_duration_amount":0},"callback":{"config":{"method":"POST","base_url":"https://api.simasinsurtech.com/","endpoint":"dataservice_test/klaim_asuransi_kredit_json.php","body_mapping":[{"operation":"shift","spec":{"LamaTelatBayar":"additional_info.overdue_payment_duration","NamaBank":"bank_name","NamaPemilikRekening":"bank_account_name","NilaiKlaim":"net_claim_amount","NoRef":"policy_details.transaction_id","NoRekeningBank":"bank_account_number","PolicyInsuranceNo":"policy_details.insurance_number"}},{"operation":"default","spec":{"SourceID":"20191030KOMUNAL"}},{"operation":"default","spec":{"IDDebitor":"01"}},{"operation":"default","spec":{"Catatan":""}}],"generated_keys":[{"name":"token","type":"auth","auth_type":"REST","auth_config":{}}],"headers":{"Authorization":"generated_keys.token.response_mapping.bearer_token","Content-Type":"application/json"},"success_response_body":{"ErrorCode":[0]}},"statuses":{"CLAIM_INITIATED":"CLAIM_INITIATED"},"status_journey":{"CLAIM_INITIATED":[]},"start_status":["CLAIM_INITIATED"]},"claim_product_validation_flag":true,"product_validation_config":{"statuses":{"CLAIM_INITIATE":true},"validation":[{"type":"RANGE","parameter":"OVERDUE_LIMIT","min":90,"max":90,"value_duration":"DAYS"}],"additional_data":null},"claim_amount_calculation_config":{"need_excess_fee":false,"balance_level":"POLICY","remaining_balance_calculation_statuses":{}}}},"partner_product_config":{"notifications":{"CLAIM_INITIATE":[],"INSURANCE_CLAIM_APPROVE":[],"INSURANCE_CLAIM_PAID":[],"INSURANCE_CLAIM_REJECT":[],"POLICY_ISSUED":[],"QOALA_CLAIM_APPROVE":[],"QOALA_CLAIM_REJECT":[]},"commissions":[{"type":"PERCENTAGE","value":0,"recipient":"PARTNER"},{"type":"PERCENTAGE","value":3,"recipient":"QOALA"}],"payment":{"payment_handler":"PAID","term_of_payment":0},"policy":{"callback":{"config":{}}},"claim":{"callback":{"config":{}}},"stoploss":{"record":false},"additional_config":{"partner_name":"PT KOMUNAL FINANSIAL INDONESIA","policy_holder":{"format":"PARTNER_NAME QQ LENDER_NAME","type":"LENDER_BENEFICIARY_NAME"}}}}},"documents":[{"type":"POLICY_CERTIFICATE","filename":"KMNL-CRDT-20220519-FN8GA.pdf","alias":"policy_certificate.pdf","url":"https://api-staging.qoala.app/api/v2/policies/documents/3d80403c-149a-49a0-a5e5-c8356bf8619e/files/KMNL-CRDT-20220519-FN8GA.pdf"}],"created_at":"2022-05-19T08:33:29Z","payment_at":"2022-05-19T08:33:29Z","commissions":null}`

	Response := `{
		"ConID": "",
		"HTTPCode": "208",
		"ErrorName": "ErrInvalidParam/ ErrParamRequired",
		"IdTransaction": "FL-52-20221017-TEST2-PSWT-20221017-002",
		"FeedbackMessage": "IDTransaction sudah pernah digunakan, untuk Nopolis : 122N0000003171 ,",
		"StatusPenerbitan": ""
	  }`

	config := schemas.UpdateDataConfig{
		TriggerStatus: map[string]string{
			"ISSUING_POLICY":                       "UUID1",
			"ISSUING_POLICY_REGEX":                 "UUID2",
			"ISSUING_POLICY_REGEX_WITH_NO_INDEX":   "UUID3",
			"ISSUING_POLICY_REGEX_NOT_FOUND":       "UUID4",
			"ISSUING_POLICY_REGEX_VALUE_NOT_FOUND": "UUID5",
		},
		Mappings: []*schemas.UpdateDataMapping{
			{
				Name:   "UUID1",
				Table:  "policies",
				Status: "success_idempotent",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "insurance_number",
						Operation: "regex[1]",
						Type:      ": (.*?) ,",
						Value:     "FeedbackMessage",
					},
				},
			},
			{
				Name:   "UUID2",
				Table:  "policies",
				Status: "success_idempotent",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "insurance_number",
						Operation: "regex",
						Type:      ": (.*?) ,",
						Value:     "FeedbackMessage",
					},
				},
			},
			{
				Name:   "UUID3",
				Table:  "policies",
				Status: "success_idempotent",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "insurance_number",
						Operation: "regex[0]",
						Type:      ": (.*?) ,",
						Value:     "FeedbackMessage",
					},
				},
			},
			{
				Name:   "UUID4",
				Table:  "policies",
				Status: "success_idempotent",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "insurance_number",
						Operation: "regex",
						Type:      ": (\\d+)) ,",
						Value:     "FeedbackMessage",
					},
					{
						Key:       "updated_at",
						Operation: "static",
						Type:      "function",
						Value:     "NOW()",
					},
				},
			},
			{
				Name:   "UUID5",
				Table:  "policies",
				Status: "success_idempotent",
				Conditions: []*schemas.ConfigKey{
					{
						Key:   "policy_number = ?",
						Value: "policy_number",
					},
				},
				ColumnMappings: []*schemas.ConfigKey{
					{
						Key:       "insurance_number",
						Operation: "regex",
						Type:      ": (.*?) ,",
						Value:     "FieldDoesNotExist",
					},
				},
			},
		},
	}

	test = append(test, []TestData{
		// Raw Queries
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "ISSUING_POLICY",
			StatusProcess:      "success_idempotent",
			Config:             config,
			Expected: []string{
				"UPDATE policies SET insurance_number = '122N0000003171' WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "ISSUING_POLICY_REGEX",
			StatusProcess:      "success_idempotent",
			Config:             config,
			Expected: []string{
				"UPDATE policies SET insurance_number = ': 122N0000003171 ,' WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "ISSUING_POLICY_REGEX_WITH_NO_INDEX",
			StatusProcess:      "success_idempotent",
			Config:             config,
			Expected: []string{
				"UPDATE policies SET insurance_number = ': 122N0000003171 ,' WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "ISSUING_POLICY_REGEX_NOT_FOUND",
			StatusProcess:      "success_idempotent",
			Config:             config,
			Expected: []string{
				// we expect the query to error out here, otherwise the execution result will be wrong
				"UPDATE policies SET , updated_at = NOW() WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
		{
			ExpectedTotalQuery: 1,
			PreparedStatement:  false,
			State:              "ISSUING_POLICY_REGEX_VALUE_NOT_FOUND",
			StatusProcess:      "success_idempotent",
			Config:             config,
			Expected: []string{
				// we expect the query to error out here, otherwise the execution result will be wrong
				"UPDATE policies SET  WHERE policy_number = 'KMNL-CRDT-20220519-FN8GA'",
			},
		},
	}...)

	t.Run("Check UpdateQueryBuilder capability", func(t *testing.T) {
		for _, v := range test {
			queries, err := UpdateQueryBuilder(v.Config, Data, Response, v.State, v.StatusProcess, v.PreparedStatement)
			if err != nil {
				assert.Fail(t, fmt.Sprintf("error occured. err: %v", err.Error()))
			}
			// Add check total query
			if v.ExpectedTotalQuery != 0 {
				assert.Equal(t, v.ExpectedTotalQuery, len(queries), "Should equal")
			}
			for k, query := range queries {
				assert.Equal(t, v.Expected[k], query, "Query should equal each other")
			}
		}
	})
}
