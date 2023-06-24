package schemas

import "time"

type PolicyBody struct {
	PolicyNumber             *string                `json:"policy_number"`
	Status                   *string                `json:"status,omitempty"`
	PartnerTransactionNumber string                 `json:"partner_transaction_number" validate:"required"`
	InsuranceNumber          *string                `json:"insurance_number"`
	TransactionNumber        string                 `json:"transaction_number" validate:"required"`
	ProductCode              string                 `json:"product_code" validate:"required"`
	ProductCategory          string                 `json:"product_category" validate:"required"`
	StartProtectionAt        *time.Time             `json:"start_protection_at" validate:"required"`
	EndProtectionAt          *time.Time             `json:"end_protection_at" validate:"required"`
	PaymentAt                *time.Time             `json:"payment_at"`
	PurchasedAt              *time.Time             `json:"purchased_at"`
	CurrencyCode             string                 `json:"currency_code"`
	PayerType                string                 `json:"payer_type"`
	Channel                  string                 `json:"channel"`
	Calculation              PolicyBodyCalculation  `json:"calculation"`
	Commissions              []PolicyBodyCommission `json:"commissions"`
	Addons                   []PolicyBodyAddon      `json:"addons"`
	PolicyHolder             PolicyBodyPolicyHolder `json:"policy_holder"`
	Booker                   *PolicyBodyBooker      `json:"booker"`
	Insureds                 []PolicyBodyInsured    `json:"insureds"`
}

type PolicyBodyCalculation struct {
	GWP                   float64 `json:"gwp"`
	PoolAmount            float64 `json:"pool_amount"`
	NettPremium           float64 `json:"nett_premium"`
	AdminFee              float64 `json:"admin_fee"`
	HardcopyAdminFee      float64 `json:"hardcopy_admin_fee"`
	MarketingSupport      float64 `json:"marketing_support"`
	AfterMarketingSupport float64 `json:"after_marketing_support"`
	TotalCommission       float64 `json:"total_commission"`
	TotalSumInsured       float64 `json:"total_sum_insured"`
	TotalAddons           float64 `json:"total_addons"`
	TotalTax              float64 `json:"total_tax"`
}

type PolicyBodyCommission struct {
	Type            string  `json:"type"`
	IdentifierId    *uint   `json:"identifier_id"`
	IdentifierValue string  `json:"identifier_value"`
	Value           float64 `json:"value"`
	Recipient       string  `json:"recipient"`
}

type PolicyBodyAddon struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Unit       int     `json:"unit"`
	Premium    float64 `json:"premium"`
	SumInsured float64 `json:"sum_insured"`
}

type PolicyBodyPolicyHolder struct {
	IdentityNumber string               `json:"identity_number"`
	IdentityType   string               `json:"identity_type"`
	FullName       string               `json:"full_name"`
	Email          string               `json:"email"`
	PhoneNumber    string               `json:"phone_number"`
	Birthdate      string               `json:"birth_date"`
	Relationship   string               `json:"relationship"`
	Gender         string               `json:"gender"`
	Address        string               `json:"address"`
	Occupation     string               `json:"occupation"`
	Documents      []PolicyBodyDocument `json:"documents"`
}

type PolicyBodyBooker struct {
	IdentityNumber string               `json:"identity_number"`
	IdentityType   string               `json:"identity_type"`
	FullName       string               `json:"full_name"`
	Email          string               `json:"email"`
	PhoneNumber    string               `json:"phone_number"`
	Birthdate      string               `json:"birth_date"`
	Relationship   string               `json:"relationship"`
	Gender         string               `json:"gender"`
	Address        string               `json:"address"`
	Occupation     string               `json:"occupation"`
	Documents      []PolicyBodyDocument `json:"documents"`
}

type PolicyBodyDocument struct {
	Filename string `json:"filename"`
	Alias    string `json:"alias"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type PolicyBodyInsured struct {
	Type      string               `json:"type"`
	Detail    interface{}          `json:"detail"`
	Documents []PolicyBodyDocument `json:"documents"`
}

type GetPolicyResponse struct {
	TransactionNumber string                    `json:"transaction_number"`
	ProductCode       string                    `json:"product_code"`
	Policies          []GetPolicyResponsePolicy `json:"policies"`
}

type GetPolicyResponsePolicy struct {
	Number    string                      `json:"number"`
	Status    string                      `json:"status"`
	Documents []GetPolicyResponseDocument `json:"documents"`
}

type GetPolicyResponseDocument struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Alias    string `json:"alias"`
	URL      string `json:"url"`
}

type CreatePolicyWithOrderNumberBody struct {
	OrderNumber string       `json:"order_number"`
	Policies    []PolicyBody `json:"policies"`
}
