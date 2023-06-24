package schemas

type Config struct {
	InsurerConfig InsurerConfig        `json:"insurer_product_config"`
	PartnerConfig PartnerProductConfig `json:"partner_product_config"`
}

type InsurerConfig struct {
	PolicyConfig PolicyConfig `json:"policy"`
	ClaimConfig  ClaimConfig  `json:"claim"`
}

func (t InsurerConfig) Validate() error {
	if err := t.PolicyConfig.Validate(); nil != err {
		return err
	}
	if err := t.ClaimConfig.Validate(); nil != err {
		return err
	}
	return nil
}

type PartnerProductConfig struct {
	Notifications    map[string][]Notification `json:"notifications"`
	CommissionConfig []CommissionConfig        `json:"commissions"`
	PaymentConfig    PaymentConfig             `json:"payment"`
	PolicyConfig     PartnerPolicyConfig       `json:"policy,omitempty"`
	ClaimConfig      PartnerClaimConfig        `json:"claim,omitempty"`
	Actions          []map[string]interface{}  `json:"actions,omitempty"`
	StoplossConfig   map[string]interface{}    `json:"stoploss,omitempty"`
	Certificate      map[string]interface{}    `json:"certificate,omitempty"`
	AdditionalConfig map[string]interface{}    `json:"additional_config,omitempty"`
	Activation       PolicyActivation          `json:"activation"`
}

func (t PartnerProductConfig) Validate() error {
	for _, c := range t.CommissionConfig {
		if err := c.Validate(); nil != err {
			return err
		}
	}
	return nil
}

type PartnerPolicyConfig struct {
	AsyncFlow bool      `json:"async_flow,omitempty"`
	Callback  *Callback `json:"callback,omitempty"`
}

type PartnerClaimConfig struct {
	AsyncFlow bool      `json:"async_flow,omitempty"`
	Callback  *Callback `json:"callback,omitempty"`
}

type Notification struct {
	Channel    string   `json:"channel"`
	Template   string   `json:"template"`
	Recipients []string `json:"recipients"`
	Cc         []string `json:"cc"`
	Bcc        []string `json:"bcc"`
	Lang       string   `json:"lang"`
}

type CommissionConfig struct {
	Type      CommissionType      `json:"type"`
	Value     float64             `json:"value"`
	Recipient CommissionRecipient `json:"recipient"`
}

func (t CommissionConfig) IsNil() bool {
	return (CommissionConfig{}) == t
}

func (t *CommissionConfig) Validate() error {
	if t.IsNil() {
		return nil
	}
	if _, err := t.Type.GetCommissionType(); nil != err {
		return err
	}
	if _, err := t.Recipient.GetCommissionTypeRecipient(); nil != err {
		return err
	}
	return nil
}

type PolicyConfig struct {
	JourneyStart                string                   `json:"journey_start"`
	JourneyEnd                  string                   `json:"journey_end"`
	Activation                  PolicyActivation         `json:"activation"`
	Expiration                  PolicyExpiration         `json:"expiration"`
	Certificate                 PolicyCertificateWording `json:"certificate"`
	Numbering                   PolicyNumbering          `json:"numbering"`
	Callback                    *Callback                `json:"callback,omitempty"`
	PolicyProductValidationFlag bool                     `json:"policy_product_validation_flag,omitempty"`
	ProductValidationConfig     *ProductValidationConfig `json:"product_validation_config,omitempty"`
	StartProtection             PolicyStartProtection    `json:"start_protection,omitempty"` // ENUM
	TransactionBreakdown        TransactionBreakdown     `json:"transaction_breakdown"`
	IsAutoAddUserPolicy         bool                     `json:"is_auto_add_user_policy"`
}

type TransactionBreakdown struct {
	ShouldBreakdown bool   `json:"should_breakdown"`
	CoveredType     string `json:"covered_type"`
}

func (t PolicyConfig) Validate() error {
	if err := t.Expiration.Validate(); nil != err {
		return err
	}
	if err := t.Certificate.Validate(); nil != err {
		return err
	}
	if err := t.Numbering.Validate(); nil != err {
		return err
	}
	return nil
}

type PaymentConfig struct {
	PaymentHandler string         `json:"payment_handler"`
	TermOfPayment  uint           `json:"term_of_payment"`
	Detail         *PaymentDetail `json:"payment_detail"`
}

type PolicyActivation struct {
	Url    string   `json:"url"`
	Fields []string `json:"fields"`
}

type PolicyExpiration struct {
	UnitDuration            UnitDurationType `json:"unit_duration"` // ENUM
	UnitDurationAmount      uint             `json:"unit_duration_amount"`
	FallbackDateType        FallbackDateType `json:"fallback_date_type"` // ENUM
	CustomPolicyPeriodField string           `json:"custom_policy_period_field"`
	FromInsuredField        string           `json:"from_insured_field"`
}

type PolicyStartProtection struct {
	UnitDuration       UnitDurationType `json:"unit_duration"` // ENUM
	UnitDurationAmount uint             `json:"unit_duration_amount,omitempty"`
	FallbackDateType   FallbackDateType `json:"fallback_date_type"` // ENUM
	FromInsuredField   string           `json:"from_insured_field"`
}

func (t PolicyExpiration) IsNil() bool {
	return (PolicyExpiration{}) == t
}

func (t *PolicyExpiration) Validate() error {
	if t.IsNil() {
		return nil
	}
	if _, err := t.UnitDuration.GetUnitDuration(); nil != err {
		return err
	}
	if _, err := t.FallbackDateType.GetFallbackDateType(); nil != err {
		return err
	}
	return nil
}

type PolicyCertificateWording struct {
	TemplateUrl            string                    `json:"template_url"`
	AlternativeTemplateUrl string                    `json:"alternative_template_url"`
	WordingPDFDocument     string                    `json:"wording"`
	Strategy               CertificateActionStrategy `json:"strategy"`
	AdditionalData         interface{}               `json:"additional_data"`
}

func (t PolicyCertificateWording) IsNil() bool {
	return (PolicyCertificateWording{}) == t
}

func (t *PolicyCertificateWording) Validate() error {
	if t.IsNil() {
		return nil
	}
	if _, err := t.Strategy.GetCertificateActionStrategy(); nil != err {
		return err
	}
	return nil
}

type PolicyNumbering struct {
	Type      PolicyNumberingType `json:"type"`
	PadLength uint                `json:"pad_length"`
	Prefix    string              `json:"prefix"` // ENUM
}

func (t PolicyNumbering) IsNil() bool {
	return (PolicyNumbering{}) == t
}

func (t *PolicyNumbering) Validate() error {
	if t.IsNil() {
		return nil
	}
	if _, err := t.Type.GetPolicyNumberingAlgorithm(); nil != err {
		return err
	}
	return nil
}

type ClaimConfig struct {
	MaxLimitation                uint                                   `json:"max_limitation"`
	JourneyEnd                   string                                 `json:"journey_end"`
	ClaimPeriodValidationSkip    bool                                   `json:"claim_period_validation_skip"`
	ClaimDeductibleFee           ClaimDeductibleFee                     `json:"claim_deductible"`
	ClaimExpiryPeriod            ClaimExpiryPeriod                      `json:"claim_expiration"`
	Callback                     *Callback                              `json:"callback,omitempty"`
	ClaimProductValidationFlag   bool                                   `json:"claim_product_validation_flag,omitempty"`
	ProductValidationConfig      *ProductValidationConfig               `json:"product_validation_config,omitempty"`
	ClaimAdditionalFlowFlag      bool                                   `json:"claim_additional_flow_flag,omitempty"`
	ClaimAdditionalFlowConfig    map[string][]ClaimAdditionalFlowConfig `json:"claim_additional_flow_config,omitempty"`
	ClaimAmountCalculationConfig *ClaimAmountCalculationConfig          `json:"claim_amount_calculation_config,omitempty"`
	ClaimStatusValidation        map[string][]string                    `json:"claim_status_validation"`
	ClaimProcessingRecepients    []string                               `json:"claim_processing_recepients"`
	CustomerConfirmClaimAmount   bool                                   `json:"customer_confirm_claim_amount"`
	ClaimAmountRejectLimit       uint                                   `json:"claim_amount_reject_limit"`
}

func (t ClaimConfig) Validate() error {
	if err := t.ClaimExpiryPeriod.Validate(); nil != err {
		return err
	}
	if err := t.ClaimDeductibleFee.Validate(); nil != err {
		return err
	}
	return nil
}

type ClaimAmountCalculationConfig struct {
	IsNeedExcessFee                     bool             `json:"need_excess_fee"`
	BalanceLevel                        BalanceLevelType `json:"balance_level"`
	RemainingBalanceCalculationStatuses map[string]bool  `json:"remaining_balance_calculation_statuses"`
}

type ClaimAdditionalFlowConfig struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	AutoTriggerStatus string `json:"auto_trigger_status,omitempty"`
	DataUpdate        *struct {
		Claim struct {
			ServiceCenterCode string `json:"service_center_code"`
			ServiceCenterName string `json:"service_center_name"`
		} `json:"claim,omitempty"`
	} `json:"data_update,omitempty"`
}

type ProductValidationConfig struct {
	Statuses       map[string]bool     `json:"statuses,omitempty"`
	Validation     []ProductValidation `json:"validation,omitempty"`
	AdditionalData map[string]struct {
		NeedClaimHistory bool `json:"need_claim_history"`
	} `json:"additional_data"`
}

type ProductValidation struct {
	Type          string  `json:"type,omitempty"`
	Parameter     string  `json:"parameter,omitempty"`
	Min           float64 `json:"min,omitempty"`
	Max           float64 `json:"max,omitempty"`
	ValueDuration string  `json:"value_duration,omitempty"`
	Value         string  `json:"value,omitempty"`
}

type ClaimDeductibleFee struct {
	UnitFee       DeductibleType `json:"unit_fee"`
	UnitFeeAmount float64        `json:"unit_fee_amount"`
}

func (t ClaimDeductibleFee) IsNil() bool {
	return (ClaimDeductibleFee{}) == t
}

func (t *ClaimDeductibleFee) Validate() error {
	if t.IsNil() {
		return nil
	}

	if _, err := t.UnitFee.GetDeductibleType(); nil != err {
		return err
	}
	return nil
}

type BenefitConfig struct {
	FormGenerator             []byte                 `json:"-"`
	ClaimWaitingPeriod        ClaimWaitingPeriod     `json:"claim_waiting_period"`
	ClaimExpiryPeriod         ClaimExpiryPeriod      `json:"claim_expiry_period"`
	ClaimLimitation           ClaimLimitation        `json:"limitation"`
	ClaimPartnerCallbackData  map[string]interface{} `json:"claim_partner_callback_data,omitempty"`
	ClaimInsurerCallbackData  map[string]interface{} `json:"claim_insurer_callback_data,omitempty"`
	PolicyPartnerCallbackData map[string]interface{} `json:"policy_partner_callback_data,omitempty"`
	PolicyInsurerCallbackData map[string]interface{} `json:"policy_insurer_callback_data,omitempty"`
	ClaimCoverage             ClaimCoverage          `json:"claim_coverage,omitempty"`
}

func (t BenefitConfig) Validate() error {
	if err := t.ClaimWaitingPeriod.Validate(); nil != err {
		return err
	}
	if err := t.ClaimExpiryPeriod.Validate(); nil != err {
		return err
	}
	return nil
}

type ClaimCoverage struct {
	ClaimCoverageType  string  `json:"claim_coverage_type"`
	ClaimCoverageValue float64 `json:"claim_coverage_value"`
	ClaimInstantPayout bool    `json:"claim_instant_payout"`
}

type ClaimWaitingPeriod struct {
	UnitDuration       UnitDurationType `json:"unit_duration"`
	UnitDurationAmount uint             `json:"unit_duration_amount"` // ENUM
}

func (t ClaimWaitingPeriod) IsNil() bool {
	return (ClaimWaitingPeriod{}) == t
}

func (t *ClaimWaitingPeriod) Validate() error {
	if t.IsNil() {
		return nil
	}
	if _, err := t.UnitDuration.GetUnitDuration(); nil != err {
		return err
	}
	return nil
}

type ClaimExpiryPeriod struct {
	UnitDuration       UnitDurationType `json:"unit_duration"`
	UnitDurationAmount uint             `json:"unit_duration_amount"` // ENUM
}

func (t ClaimExpiryPeriod) IsNil() bool {
	return (ClaimExpiryPeriod{}) == t
}

func (t *ClaimExpiryPeriod) Validate() error {
	if t.IsNil() {
		return nil
	}
	if _, err := t.UnitDuration.GetUnitDuration(); nil != err {
		return err
	}
	return nil
}

type ClaimLimitation struct {
	MaxLimitation               uint   `json:"max_limitation"`
	SameIdentityClaimCount      uint   `json:"same_identity_claim_count"`
	SameIdentityClaimCountField string `json:"same_identity_claim_count_field"`
}

type Callback struct {
	Enable         bool                `json:"enable"`
	CallbackConfig CallbackConfig      `json:"config,omitempty"`
	Statuses       map[string]string   `json:"statuses,omitempty"`
	StatusJourney  map[string][]string `json:"status_journey,omitempty"`
	StartStatus    []string            `json:"start_status,omitempty"`
}

type CallbackConfig struct {
	Method              string                   `json:"method,omitempty"`
	BaseUrl             string                   `json:"base_url,omitempty"`
	EndPoint            string                   `json:"endpoint,omitempty"`
	BodyMapping         []interface{}            `json:"body_mapping,omitempty"`
	StaticBodyMapping   map[string]interface{}   `json:"static_body_mapping,omitempty"`
	UpdateBodyMapping   []*updateBodyMappingKey  `json:"update_body_mapping"`
	ConfigKeys          []*ConfigKey             `json:"config_keys,omitempty"`
	GeneratedKeys       []*GeneratedKey          `json:"generated_keys,omitempty"`
	Headers             map[string]string        `json:"headers,omitempty"`
	ResponseMapping     map[string]string        `json:"response_mapping,omitempty"`
	SuccessResponseBody map[string][]interface{} `json:"success_response_body,omitempty"`
	UpdateDataConfig    *UpdateDataConfig        `json:"update_data_config,omitempty"`
}

type ConfigKey struct {
	Name      string `json:"name,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Type      string `json:"type,omitempty"`
	Operation string `json:"operation,omitempty"`
}

type updateBodyMappingKey struct {
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Type      string `json:"type,omitempty"`
	Operation string `json:"operation,omitempty"`
}

type GeneratedKey struct {
	Name        string                  `json:"name,omitempty"`
	Key         string                  `json:"key,omitempty"`
	Format      string                  `json:"format,omitempty"`
	DataMapping string                  `json:"data_mapping,omitempty"`
	Type        string                  `json:"type"`
	AuthType    string                  `json:"auth_type,omitempty"`
	AuthConfig  *GeneratedKeyAuthConfig `json:"auth_config,omitempty"`
}

type GeneratedKeyAuthConfig struct {
	Secret            string                       `json:"secret,omitempty"`
	Encoding          string                       `json:"encoding,omitempty"`
	MessageGeneration *AuthConfigMessageGeneration `json:"message_generation,omitempty"`
	RestConfig        *CallbackConfig              `json:"rest_config,omitempty"`
}

type AuthConfigMessageGeneration struct {
	Format string   `json:"format,omitempty"`
	Params []string `json:"params,omitempty"`
}

type CallbackConfigData struct {
	PolicyCallback CallbackDetail `json:"policy"`
	ClaimCallback  CallbackDetail `json:"claim"`
}

type CallbackDetail struct {
	Callback Callback `json:"callback,omitempty"`
}

type UpdateDataConfig struct {
	TriggerStatus map[string]string    `json:"trigger_status,omitempty"`
	Mappings      []*UpdateDataMapping `json:"mappings,omitempty"`
}

type UpdateDataMapping struct {
	Name           string       `json:"name,omitempty"`
	Table          string       `json:"table,omitempty"`
	Status         string       `json:"status,omitempty"`
	Conditions     []*ConfigKey `json:"conditions,omitempty"`
	ColumnMappings []*ConfigKey `json:"column_mappings,omitempty"`
}

type CallbackDataUpdateRequest struct {
	Config                 CallbackDataUpdateConfig `json:"config,omitempty"`
	ClaimNumber            string                   `json:"claim_number,omitempty"`
	ClaimStatus            string                   `json:"claim_status,omitempty"`
	PolicyNumber           string                   `json:"policy_number,omitempty"`
	PolicyStatus           string                   `json:"policy_status,omitempty"`
	InsurerPartnerResponse interface{}              `json:"insurer_partner_response,omitempty"`
	InsurerPartnerStatus   string                   `json:"insurer_partner_status,omitempty"`
	Queries                []string                 `json:"queries,omitempty"`
	PublishUnixTimestamp   int64                    `json:"publish_unix_timestamp"`
	Retries                int64                    `json:"retries"`
	LatestErrorMessage     *string                  `json:"latest_error_message,omitempty"`
	IsManualHit            *bool                    `json:"is_manual_hit,omitempty"`
}

type CallbackDataUpdateConfig struct {
	UpdateDataConfig UpdateDataConfig `json:"update_data_config,omitempty"`
}

type FlightDelayActionConfig struct {
	Enabled        bool           `json:"enabled"`
	DelayThreshold int            `json:"delay_threshold"`
	Provider       string         `json:"provider"`
	Notifications  []Notification `json:"notifications"`
}
