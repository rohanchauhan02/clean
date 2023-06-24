package schemas

import "fmt"

type InsurerConfigDTO struct {
	Alternative string `json:"alternative"`
	Config      InsurerConfig
}

type InsuranceProductOnBoardingDTO struct {
	Name                 string             `json:"name"`
	ProductCode          string             `json:"product_code"`
	InsuranceProductCode string             `json:"insurance_product_code"`
	Summary              string             `json:"summary"`
	Description          string             `json:"description"`
	ClaimDescription     string             `json:"claim_description"`
	Term                 string             `json:"term"`
	CategoryCode         string             `json:"category_code"`
	Currency             string             `json:"currency"`
	Country              string             `json:"country"`
	PremiumType          PremiumType        `json:"premium_type"`
	AdminFee             float64            `json:"admin_fee"`
	HardCopyAdminFee     float64            `json:"hardcopy_admin_fee"`
	MarketingSupportFee  float64            `json:"marketing_support_fee"`
	TaxFee               float64            `json:"tax_fee"`
	StampFee             float64            `json:"stamp_fee"`
	InsurerConfig        InsurerConfig      `json:"insurer_config"`
	CallbackConfig       CallbackConfigData `json:"callback_config"`
	DurationValue        int                `json:"duration_value"`
	DurationType         UnitDurationType   `json:"duration_type"`
	Language             string             `json:"language"`
	Timezone             string             `json:"timezone"`
	InsuranceCode        string             `json:"insurance_code"`
	Rules                interface{}        `json:"rules"`
	Benefits             []Benefit          `json:"benefits"`
	IsActive             bool               `json:"is_active"`
	IsRenewalProduct     bool               `json:"is_renewal_product"`
	IsSoftDelete         bool               `json:"is_soft_delete"`
}

func (t *InsuranceProductOnBoardingDTO) Validate() error {
	if _, err := t.PremiumType.GetPremiumType(); nil != err {
		return err
	}
	if _, err := t.DurationType.GetUnitDuration(); nil != err {
		return err
	}
	if err := t.InsurerConfig.Validate(); nil != err {
		return err
	}
	for _, obj := range t.Benefits {
		if err := obj.Validate(); nil != err {
			return err
		}
	}

	if t.InsuranceCode == "" {
		return fmt.Errorf("insurance code cannot be empty")
	}

	return nil
}

type InsuranceProductOnBoardingResponse struct {
	Id                   string             `json:"id"`
	InsuranceProductCode string             `json:"insurance_product_code"`
	Name                 string             `json:"name"`
	Summary              string             `json:"summary"`
	Description          string             `json:"description"`
	ClaimDescription     string             `json:"claim_description"`
	Term                 string             `json:"term"`
	CategoryCode         string             `json:"category_code"`
	CategoryName         string             `json:"category_name"`
	Currency             string             `json:"currency"`
	Country              string             `json:"country"`
	PremiumType          PremiumType        `json:"premium_type"`
	AdminFee             float64            `json:"admin_fee"`
	HardCopyAdminFee     float64            `json:"hardcopy_admin_fee"`
	MarketingSupportFee  float64            `json:"marketing_support_fee"`
	TaxFee               float64            `json:"tax_fee"`
	StampFee             float64            `json:"stamp_fee"`
	InsurerConfig        InsurerConfig      `json:"insurer_config"`
	CallbackConfig       CallbackConfigData `json:"callback_config"`
	DurationValue        int                `json:"duration_value"`
	DurationType         UnitDurationType   `json:"duration_type"`
	Language             string             `json:"language"`
	Timezone             string             `json:"timezone"`
	InsuranceCode        string             `json:"insurance_code"`
	Rules                interface{}        `json:"rules"`
	Benefits             []Benefit          `json:"benefits"`
}

type PartnerProductOnBoardingDTO struct {
	ProductId            string               `json:"product_id"`
	InsuranceProductCode string               `json:"insurance_product_code"`
	PartnerCode          string               `json:"partner_code"`
	ProductName          string               `json:"product_name"`
	ProductCode          string               `json:"product_code"`
	Config               PartnerProductConfig `json:"config"`
	CallbackConfig       CallbackConfigData   `json:"callback_config"`
	IsActive             bool                 `json:"is_active"`
	IsSoftDelete         bool                 `json:"is_soft_delete"`
}

func (t PartnerProductOnBoardingDTO) Validate() error {
	if err := t.Config.Validate(); err != nil {
		return err
	}

	if t.ProductId == "" && t.InsuranceProductCode == "" {
		return fmt.Errorf("insurance product id or product code cannot be empty")
	}

	if t.PartnerCode == "" {
		return fmt.Errorf("partner code cannot be empty")
	}
	return nil
}

type PartnerProductOnBoardingResponse struct {
	Id                   string               `json:"id"`
	ProductId            string               `json:"product_id"`
	PartnerCode          string               `json:"partner_code"`
	ProductName          string               `json:"product_name"`
	ProductCode          string               `json:"product_code"`
	InsuranceProductCode string               `json:"insurance_product_code"`
	Config               PartnerProductConfig `json:"config"`
	CallbackConfig       CallbackConfigData   `json:"callback_config"`
}
