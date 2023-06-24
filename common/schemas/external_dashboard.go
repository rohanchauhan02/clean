package schemas

import (
	"encoding/json"
	"time"
)

type (
	ClaimDetailSchemaClaimHistory struct {
		ID          int64      `json:"id"`
		Status      string     `json:"status"`
		Description string     `json:"description"`
		CreatedAt   time.Time  `json:"created_at"`
		CreatedBy   string     `json:"created_by"`
		UpdatedAt   *time.Time `json:"updated_at"`
	}

	ClaimDetailRequestDTO struct {
		UserType                     string                          `json:"user_type,required"`
		ClaimStatus                  string                          `json:"claim_status,required"`
		ProductCategory              string                          `json:"product_category,required"`
		ProductCode                  string                          `json:"product_code,required"`
		BenefitCode                  string                          `json:"benefit_code,required"`
		ClaimHistories               []ClaimDetailSchemaClaimHistory `json:"claim_histories"`
		ClaimAdditionalInfo          json.RawMessage                 `json:"claim_additional_info,omitempty"`
		TotalAmount                  float64                         `json:"total_amount,omitempty"`
		TotalCovered                 int                             `json:"total_covered,omitempty"`
		RequestedDocuments           json.RawMessage                 `json:"requested_documents,omitempty"`
		RequestedAdditionalDocuments json.RawMessage                 `json:"requested_additional_documents,omitempty"`
	}

	ClaimListRequestDTO struct {
		State           string               `json:"state,omitempty"`
		UserType        string               `json:"user_type,omitempty"`
		ProductCategory *ProductCategoryCode `json:"product_category,omitempty"`
		InsurerCode     string               `json:"insurer_code,omitempty"`
	}

	// Provide standard constant for available product category
	StandardProductCategory struct {
		Code  string `json:"code"`
		Name  string `json:"name"`
		Alias string `json:"alias"`
	}

	ProductDashboardListResponse struct {
		Name                 string    `json:"name"`
		Description          string    `json:"description"`
		InsurancePartnerCode string    `json:"insurance_partner_code,omitempty"`
		PartnerCode          string    `json:"partner_code,omitempty"`
		Code                 string    `json:"code"`
		IsActive             bool      `json:"is_active"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
	}

	SoftDeleteProductResponse struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		ProductType string    `json:"product_type"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		DeletedAt   time.Time `json:"deleted_at"`
	}

	EditStatusProductResponse struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		ProductType string    `json:"product_type"`
		IsActive    bool      `json:"is_active"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

// product category to parent category on dashboard mapping
var ProductCategoryParentCategoryMap = map[string]string{
	"PA+":               "microhealth",
	"CRDL":              "finance",
	"CRDT":              "finance",
	"SCREEN_PROTECTION": "gadget",
	"FULL_PROTECTION":   "gadget",
}

var ParentCategoryToProductCategories = map[string][]string{
	"microhealth": {
		"PA+",
	},
	"finance": {
		"CRDL",
		"CRDT",
		"CRDC",
		"CRDP",
	},
	"gadget": {
		"SCREEN_PROTECTION",
		"FULL_PROTECTION",
	},
	"travel": {
		"FLIGHT",
		"TRIN",
		"HTEL",
		"TBUS",
		"EXPR",
	},
}

var StandardProductCategories = `
{
  "FINANCE": [
    {
      "alias": "ALL",
      "code": "ALL",
      "name": "All"
    },
    {
      "alias": "CRDT",
      "code": "CRDT",
      "name": "Credit Insurance with Reinstatement"
    },
    {
      "alias": "CRDL",
      "code": "CRDL",
      "name": "Credit Life Insurance"
    },
    {
      "alias": "CRDP",
      "code": "CRDP",
      "name": "Credit Insurance for Productive Loan"
    },
    {
      "alias": "CRDC",
      "code": "CRDC",
      "name": "Credit Insurance for Consumptive Loan"
    }
  ],
  "MICROHEALTH": [
    {
      "alias": "ALL",
      "code": "ALL",
      "name": "All"
    },
    {
      "alias": "MHPA",
      "code": "PA+",
      "name": "Personal Accident+"
    },
    {
      "alias": "MHPA",
      "code": "PA",
      "name": "Personal Accident"
    },
    {
      "code": "DBD",
      "name": "Dengue Fever"
    },
    {
      "alias": "MHCP",
      "code": "HCP",
      "name": "Hospital Cash Plan"
    }
  ],
  "GADGET": [
    {
      "code": "ALL",
      "name": "All"
    },
    {
      "code": "SCREEN_PROTECTION",
      "name": "Screen Protection"
    },
    {
      "code": "FULL_PROTECTION",
      "name": "Full Protection"
    }
  ],
  "TRAVEL": [
    {
      "code": "ALL",
      "name": "All"
    },
    {
      "code": "FLIGHT",
      "name": "Flight"
    },
    {
      "code": "TRIN",
      "name": "Train"
    },
    {
      "code": "HTEL",
      "name": "Hotel"
    },
    {
      "code": "TBUS",
      "name": "Bus"
    },
    {
      "code": "EXPR",
      "name": "Experience"
    }
  ]
}`

func GetStandardProductCategories(parentCategory string) []StandardProductCategory {
	mappedResp := []StandardProductCategory{}

	stdParentProductCategories := map[string][]StandardProductCategory{}
	err := json.Unmarshal([]byte(StandardProductCategories), &stdParentProductCategories)
	if err != nil {
		return mappedResp
	}

	mappedResp = stdParentProductCategories[parentCategory]
	return mappedResp
}
