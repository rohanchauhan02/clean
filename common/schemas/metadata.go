package schemas

import (
	"encoding/json"
	"time"
)

type Metadata struct {
	Id         string          `json:"id"`
	Kind       string          `json:"kind"`
	ParentCode string          `json:"parentCode"`
	Code       string          `json:"code"`
	Level      int             `json:"level"`
	Name       string          `json:"name"`
	Additional json.RawMessage `json:"additional"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}

type MetadataAdditional struct {
	Location struct {
		Level1    string  `json:"level_1"`
		Level2    string  `json:"level_2"`
		Level3    string  `json:"level_3"`
		Timezone  string  `json:"timezone"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	ImageUrl string `json:"image_url"`
}
