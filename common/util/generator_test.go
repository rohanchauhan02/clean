package util

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dataStruct struct {
	MainCategory    string `json:"main_category"`
	ProductCategory string `json:"product_category"`
	PartnerCode     string `json:"partner_code"`
	LatestValue     int64  `json:"latest_value"`
}

func TestGenerateBenefitCode(t *testing.T) {
	t.Run("test generate code with format", func(t *testing.T) {
		code := BenefitCodeGenerator("PA", "Takaful", "Accidental Death Benefit")
		assert.Contains(t, code, "PA-TAKAFUL-ADB")
	})
}

func TestProductCodeGenerator(t *testing.T) {
	t.Run("test generate productCode with format", func(t *testing.T) {
		var TestingData []dataStruct
		data := `[{"main_category":"FINANCE","product_category":"CRDT","partner_code":"KMNL","latest_value":1},{"main_category":"MICROHEALTH","product_category":"MHCP","partner_code":"FNMS","latest_value":4},{"main_category":"TRAVEL","product_category":"FLGT","partner_code":"TVLK","latest_value":0},{"main_category":"GADGET","product_category":"MR12","partner_code":"TKPD","latest_value":2}]`
		err := json.Unmarshal([]byte(data), &TestingData)
		if err != nil {
			t.Fatalf("error in unmarshalling json string for mock data with error: %s", err.Error())
		}
		for _, value := range TestingData {
			code := ProductCodeGenerator(value.MainCategory, value.ProductCategory, value.PartnerCode, value.LatestValue)
		testingBreak:
			switch value.MainCategory {
			case "FINANCE":
				assert.Contains(t, code, "F-CRDT-KMNL-002")
				break testingBreak
			case "MICROHEALTH":
				assert.Contains(t, code, "M-MHCP-FNMS-005")
				break testingBreak
			case "TRAVEL":
				assert.Contains(t, code, "T-FLGT-TVLK-001")
				break testingBreak
			case "GADGET":
				assert.Contains(t, code, "G-MR12-TKPD-003")
				break testingBreak
			}
		}
	})
}
