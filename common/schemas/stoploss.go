package schemas

import "time"

type StoplossDetail struct {
	Code                string         `json:"code"`
	PartnerCode         string         `json:"partner_code"`
	InsuranceCode       string         `json:"insurance_code"`
	ClaimedPremium      float64        `json:"claimed_premium"`
	TotalPremium        float64        `json:"total_premium"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	ProductsCode        []string       `json:"products_code,omitempty"`
	StoplossConfig      *PaymentDetail `json:"config,omitempty"`
	DisbursementEnabled bool           `json:"disbursement_enabled"`
	InvoiceEnabled      bool           `json:"invoice_enabled"`
	Insurance           Insurance      `json:"insurance,omitempty"`
	Partner             Partner        `json:"partner,omitempty"`
}
