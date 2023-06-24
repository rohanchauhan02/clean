package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rohanchauhan02/clean/common/schemas"
)

func BenefitCodeGenerator(productCategoryCode string, insurerCode string, benefitName string) string {
	CodeFormat := "%s-%s-%s-%s"

	benefitPart := strings.Split(benefitName, " ")
	benefitName = ""
	for _, p := range benefitPart {
		benefitName += string(p[0])
	}

	return strings.ToUpper(fmt.Sprintf(CodeFormat, productCategoryCode, insurerCode, benefitName, RandomString(5, "UPPERCASE_ALPHANUMERIC")))
}

func ProductCodeGenerator(mainCategory string, benefitCode string, codeAlias string, latestValue int64) string {
	StandardProductCategories := schemas.StandardProductCategories
	stdParentProductCategories := map[string][]schemas.StandardProductCategory{}
	CodeFormat := "%s-%s-%s-%s"

	err := json.Unmarshal([]byte(StandardProductCategories), &stdParentProductCategories)
	if err != nil {
		return fmt.Sprintf("failed UnmarshalData with Error: %s", err.Error())
	}

	category := stdParentProductCategories[strings.ToUpper(mainCategory)]
	benefitAlias := benefitCode
	for _, value := range category {
		if value.Code == benefitCode && value.Alias != "" {
			benefitAlias = value.Alias
		}
	}

	latestValue++
	newIncrement := fmt.Sprintf("%03d", latestValue)

	return strings.ToUpper(fmt.Sprintf(CodeFormat, string(mainCategory[0]), benefitAlias, codeAlias, newIncrement))
}
