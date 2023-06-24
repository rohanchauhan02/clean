package models

import (
	"time"

	"github.com/rohanchauhan02/clean/common/schemas"
)

// ResponsePattern return object  general response pattern
type ResponsePattern struct {
	Status    string               `json:"status"`
	Data      interface{}          `json:"data"`
	Message   string               `json:"message"`
	Code      int                  `json:"code"`
	Meta      *ResponsePatternMeta `json:"meta,omitempty"`
	ErrorCode string               `json:"error_code,omitempty"`
}

// ResponsePatternMeta return object for response meta
type ResponsePatternMeta struct {
	Limit *int `json:"limit,omitempty"`
	Page  *int `json:"page,omitempty"`
	Total *int `json:"total,omitempty"`
	Count *int `json:"count,omitempty"`
}

// ResponseInternalPattern return internal object general response pattern
type ResponseInternalPattern struct {
	Status    string                       `json:"status"`
	Data      interface{}                  `json:"data"`
	Message   string                       `json:"message"`
	Code      int                          `json:"code"`
	Meta      *ResponseInternalPatternMeta `json:"meta,omitempty"`
	ErrorCode string                       `json:"error_code,omitempty"`
}

// ResponseInternalPatternMeta return object for Qoala internal response meta
type ResponseInternalPatternMeta struct {
	Count            int                    `json:"count"`
	Page             int                    `json:"page"`
	Limit            int                    `json:"limit"`
	Pages            int                    `json:"pages"`
	Select           []string               `json:"select"`
	AvailableSelects []MetaAvailableSelects `json:"availableSelects"`
	Statuses         []MetaStatuses         `json:"statuses"`
	States           []MetaStates           `json:"states"`
}

type MetaAvailableSelects struct {
	Field          string      `json:"field"`
	FallbackField  string      `json:"fallbackField"`
	Type           string      `json:"type"`
	Label          string      `json:"label"`
	CurrencyField  string      `json:"currencyField"`
	IsSortable     bool        `json:"isSortable"`
	IsSelectable   bool        `json:"IsSelectable"`
	IsShowByHidden bool        `json:"isShowByHidden"`
	RestrictTypes  []string    `json:"restrict_types"`
	DefaultValue   interface{} `json:"default_value"`
}

type MetaStatuses struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MetaStates struct {
	Label string `json:"_label"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ClaimProductValidityRequest struct {
	Number          string                              `json:"number"`
	LastStatus      string                              `json:"last_status"`
	TotalAmount     float64                             `json:"total_amount"`
	ExcessFee       float64                             `json:"excess_fee"`
	DeductibleFee   *float64                            `json:"deductible_fee"`
	ClaimBalances   ClaimProductValidityClaimBalances   `json:"claim_balances"`
	CreatedAt       time.Time                           `json:"created_at"`
	PolicyDetails   ClaimProductValidityPolicyDetails   `json:"policy_details"`
	ProductDetails  ClaimProductValidityProductDetails  `json:"product_details"`
	ClaimHistories  []ClaimProductValidityClaimHistory  `json:"claim_histories"`
	AdditionalInfo  map[string]interface{}              `json:"additional_info"`
	UpdateClaimData ClaimProductValidityUpdateClaimData `json:"update_claim_data"`
	CustomerTravel  ClaimProductValidityCustomerTravel  `json:"customer_travel"`
	TotalCovered    int                                 `json:"total_covered,omitempty"`
	CustomerLoan    schemas.LoanInsuredDetail           `json:"customer_loan"`
}

type ClaimProductValidityCustomerTravel struct {
	AdditionalInfo map[string]interface{} `json:"additional_info"`
}

type ClaimProductValidityClaimBalances struct {
	SumInsured float64 `json:"sum_insured"`
}

type ClaimProductValidityPolicyDetails struct {
	InsuranceCode     string    `json:"insurance_code"`
	StartProtectionAt time.Time `json:"start_protection_at"`
	EndProtectionAt   time.Time `json:"end_protection_at"`
	SumInsured        float64   `json:"sum_insured"`
	RemainingBalance  *float64  `json:"remaining_balance"`
}

type ClaimProductValidityProductDetails struct {
	Code        string `json:"code"`
	BenefitCode string `json:"benefit_code"`
}

type ClaimProductValidityResponse struct {
	Status      string                           `json:"status"`
	Message     string                           `json:"message"`
	DataUpdate  *ClaimProductValidityDataUpdate  `json:"data_update,omitempty"`
	AutoTrigger *ClaimProductValidityAutoTrigger `json:"auto_trigger,omitempty"`
}

type ClaimProductValidityDataUpdate struct {
	Claim *struct {
		TotalAmount   *float64 `json:"total_amount,omitempty"`
		LastStatus    *string  `json:"last_status,omitempty"`
		DeductibleFee *float64 `json:"deductible_fee,omitempty"`
		ProductCode   *string  `json:"product_code,omitempty"`
	} `json:"claim,omitempty"`
	ClaimHistory *struct {
		Status      *string `json:"status,omitempty"`
		Description *string `json:"description,omitempty"`
	} `json:"claim_history,omitempty"`
	ClaimBalance *struct {
		SumInsured *float64 `json:"sum_insured,omitempty"`
	} `json:"claim_balance,omitempty"`
	ClaimAdditionalInfo map[string]interface{} `json:"claim_additional_info,omitempty"`
}

type ClaimProductValidityRequestedAction struct {
	Type               string `json:"type"`
	Title              string `json:"title"`
	InternalTitle      string `json:"_title"`
	DetailValueMapping []struct {
		InternalLabel string `json:"_label"`
		Label         string `json:"label"`
		Value         string `json:"value"`
	} `json:"detail_value_mapping"`
	OptionConfig interface{} `json:"option_config"`
}

type ClaimProductValidityAutoTrigger struct {
	Status          *string                              `json:"status,omitempty"`
	RequestedAction *ClaimProductValidityRequestedAction `json:"requested_action,omitempty"`
}

type ClaimProductValidityClaimHistory struct {
	ID          int64      `json:"id"`
	Status      string     `json:"status"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CreatedBy   string     `json:"created_by"`
}

type ClaimProductValidityUpdateClaimData struct {
	Amount      float64 `json:"amount"`
	Action      string  `json:"action"`
	Note        string  `json:"note"`
	ProductCode string  `json:"product_code"`
}

// ResponseGeneralPattern return object general response pattern with status int
type ResponseGeneralPattern struct {
	Status    int                  `json:"status"`
	Data      interface{}          `json:"data"`
	Message   string               `json:"message"`
	Code      int                  `json:"code"`
	Meta      *ResponsePatternMeta `json:"meta,omitempty"`
	ErrorCode string               `json:"error_code,omitempty"`
}
