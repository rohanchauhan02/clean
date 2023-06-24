// These are helper structs that helps gives more verbosity to the
// DTO for ProductOnboarding as well as ProductList and ProductDetail
// APIs that are needed

package schemas

import (
	"gorm.io/datatypes"

	"time"
)

// Benefit defines the DTO object that's usually used
// for product onboarding and also displayed in
// ProductList APIs
type Benefit struct {
	Id                      string         `json:"id,omitempty"`
	Name                    string         `json:"name"`
	Logo                    string         `json:"logo"`
	Code                    string         `json:"code"`
	Amount                  float64        `json:"amount"`
	Type                    string         `json:"type"`
	Description             string         `json:"description"`
	ClaimRequiredDocuments  []string       `json:"claim_required_documents,omitempty"`
	ClaimSupportedDocuments []string       `json:"claim_supported_documents,omitempty"`
	BenefitConfig           *BenefitConfig `json:"config,omitempty"`
	FormConfig              datatypes.JSON `json:"form_config,omitempty"`
	IsSoftDelete            bool           `json:"is_soft_delete,omitempty"`
}

func (t Benefit) Validate() error {
	if err := t.BenefitConfig.Validate(); nil != err {
		return err
	}
	return nil
}

// Insurance defines Insurance object sent from External Dashboard
// during ProductOnboarding
type Insurance struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Logo        string `json:"logo"`
	LogoSquare  string `json:"logo_square"`
	Description string `json:"description"`
}

// Partner defines Partner object sent from External Dashboard
// during ProductOnboarding
type Partner struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Logo        string `json:"logo"`
	LogoSquare  string `json:"logo_square"`
	Description string `json:"description"`
}

type ProductDetailResponse struct {
	Id                        string      `json:"id"`
	ProductCode               string      `json:"product_code"`
	InsurerProductName        string      `json:"insurer_product_name"`
	InsurerProductCode        string      `json:"insurer_product_code"`
	ProductName               string      `json:"product_name"`
	Country                   string      `json:"country"`
	Currency                  string      `json:"currency"`
	ProductSummary            string      `json:"product_summary"`
	ProductDescription        string      `json:"product_description"`
	ProductTerms              string      `json:"product_terms"`
	ClaimDescription          string      `json:"claim_description"`
	IsRenewalProduct          bool        `json:"is_renewal_product"`
	Language                  string      `json:"language"`
	Timezone                  string      `json:"timezone"`
	ParentCategory            string      `json:"parent_category"`
	IsActive                  bool        `json:"is_active"`
	CategoryName              string      `json:"category_name"`
	CategoryCode              string      `json:"category_code"`
	AdminFee                  float64     `json:"admin_fee"`
	HardCopyAdminFee          float64     `json:"hardcopy_admin_fee"`
	MarketingSupportFee       float64     `json:"marketing_support_fee"`
	TaxFee                    float64     `json:"tax_fee"`
	StampFee                  float64     `json:"stamp_fee"`
	PremiumType               string      `json:"premium_type"`
	Premium                   []float64   `json:"premium"`
	NettPremium               []float64   `json:"nett_premium"`
	SumInsured                []float64   `json:"sum_insured"`
	PoolAmount                []float64   `json:"pool_amount,omitempty"`
	PremiumUnitDuration       string      `json:"premium_unit_duration"`
	PremiumUnitDurationAmount float64     `json:"premium_unit_duration_amount"`
	CreatedAt                 time.Time   `json:"created_at,omitempty"`
	UpdatedAt                 time.Time   `json:"updated_at,omitempty"`
	Insurance                 Insurance   `json:"insurance"`
	Partner                   Partner     `json:"partner"`
	Benefits                  []Benefit   `json:"benefits"`
	Config                    *Config     `json:"config,omitempty"`
	AdditionalInfo            interface{} `json:"additional_info"`
}

// CategoryMap maps category name for each category code used
// in the dashboard
var CategoryMap map[string]string = map[string]string{
	"CRDT":              "Credit Insurance with Reinstatement",
	"CLI":               "Credit Life Insurance",
	"CRDP":              "Credit Insurance for Productive Loan",
	"CRDC":              "Credit Insurance for Consumptive Loan",
	"SCREEN_PROTECTION": "Screen Protection",
	"PA+":               "Personal Accident+",
	"PA":                "Personal Accident",
	"HCP":               "Hospital Cash Plan",
	"DBD":               "Dengue Fever",
	"HTEL":              "Hotel Insurance",
	"FLIGHT":            "Flight Insurance",
	"EXPR":              "Experience Insurance",
	"TRIN":              "Train Insurance",
	"TBUS":              "Bus Insurance",
}

type QuotationUpdateStatusQuote struct {
	PartnerTransactionId string `json:"partner_transaction_id"`
	Status               string `json:"status"`
}

type FinanceUpdateStoploss struct {
	PartnerTransactionId string    `json:"partner_transaction_id,omitempty"`
	PolicyNumber         string    `json:"policy_number,omitempty"`
	ClaimNumber          string    `json:"claim_number,omitempty"`
	PartnerProductCode   string    `json:"partner_product_code,omitempty"`
	Premium              float64   `json:"premium,omitempty"`
	Status               string    `json:"status,omitempty"`
	UserName             string    `json:"user_name,omitempty"`
	InsertHistory        bool      `json:"insert_history"`
	StoplossCode         string    `json:"stoploss_code"`
	TransactionType      string    `json:"transaction_type"`
	TransactionDate      time.Time `json:"transaction_date"`
	IdentifierType       string    `json:"identifier_type"`
	IdentifierNumber     string    `json:"identifier_number"`
	Amount               float64   `json:"amount,omitempty"`
	UnroundedAmount      float64   `json:"unrounded_amount,omitempty"`
	PoolAmount           float64   `json:"pool_amount,omitempty"`
	UnroundedPoolAmount  float64   `json:"unrounded_pool_amount,omitempty"`
}
