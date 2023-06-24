package schemas

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStandardProductCategories(t *testing.T) {
	t.Run("Test get microhealth parent category", func(t *testing.T) {
		parentCategory := "MICROHEALTH"

		stdParentProductCategories := map[string][]StandardProductCategory{}
		err := json.Unmarshal([]byte(StandardProductCategories), &stdParentProductCategories)
		assert.Nil(t, err, "test ok unmarshal expected resp")
		expectedResp := stdParentProductCategories[parentCategory]

		actualResp := GetStandardProductCategories(parentCategory)
		assert.Equal(t, expectedResp, actualResp, "success get microheath product categories")
	})

	t.Run("Test get finance parent category", func(t *testing.T) {
		parentCategory := "FINANCE"

		stdParentProductCategories := map[string][]StandardProductCategory{}
		err := json.Unmarshal([]byte(StandardProductCategories), &stdParentProductCategories)
		assert.Nil(t, err, "test ok unmarshal expected resp")
		expectedResp := stdParentProductCategories[parentCategory]

		actualResp := GetStandardProductCategories(parentCategory)
		assert.Equal(t, expectedResp, actualResp, "success get finance product categories")
	})

	t.Run("Test get smartphone parent category", func(t *testing.T) {
		parentCategory := "GADGET"

		stdParentProductCategories := map[string][]StandardProductCategory{}
		err := json.Unmarshal([]byte(StandardProductCategories), &stdParentProductCategories)
		assert.Nil(t, err, "test ok unmarshal expected resp")
		expectedResp := stdParentProductCategories[parentCategory]

		actualResp := GetStandardProductCategories(parentCategory)
		assert.Equal(t, expectedResp, actualResp, "success get smartphone product categories")
	})
}
