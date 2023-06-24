package schemas

type UserServiceOrganizationsResponse struct {
	Data []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		Code        string      `json:"code"`
		CodeAlias   string      `json:"codeAlias"`
		Country     string      `json:"country"`
		PartnerID   int         `json:"partnerId"`
		Alternative string      `json:"alternative"`
		Description string      `json:"description"`
		Type        string      `json:"type"`
		Logo        string      `json:"logo"`
		LogoSmall   interface{} `json:"logoSmall"`
		ColorTheme  struct {
			Primary   string `json:"primary"`
			Secondary string `json:"secondary"`
		} `json:"colorTheme"`
		Config struct {
			HasLevel      bool `json:"hasLevel"`
			HasAgent      bool `json:"hasAgent"`
			HasCommission bool `json:"hasCommission"`
		} `json:"config"`
		Contact   interface{} `json:"contact"`
		Address   interface{} `json:"address"`
		DeletedAt interface{} `json:"deletedAt"`
		Insurance struct {
			ID    int    `json:"id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			Alias string `json:"alias"`
		} `json:"insurance"`
		ProductCategories []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Industry    string `json:"industry"`
			Icon        string `json:"icon"`
			IsDraft     int    `json:"isDraft"`
		} `json:"productCategories"`
		Industries []string      `json:"industries"`
		Relations  []interface{} `json:"relations"`
		Bank       interface{}   `json:"bank"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Pages int `json:"pages"`
	} `json:"meta"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
