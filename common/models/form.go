package models

type (
	FormConfigData struct {
		DataConfig       FormDataConfig       `json:"data_config"`
		ValidationConfig FormValidationConfig `json:"validation_config"`
	}

	FormDataConfig struct {
		DocumentObjects []FormConfigDocumentObject `json:"document_objects"`
	}

	FormConfigDocumentObject struct {
		Key               string `json:"key"`
		EndUserVisibility bool   `json:"end_user_visibility"`
	}

	FormValidationConfig struct {
		RequiredFields []string `json:"required_fields"`
	}
)
