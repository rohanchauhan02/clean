package country_validation

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	CountryCode                       = "CountryCode"
	pluginName                        = "gorm-country-validation"
	customGormEventCreateCallbackName = pluginName + ":before_create"
	customGormEventUpdateCallbackName = pluginName + ":before_update"

	gormEventCreateName = "gorm:create"
	gormEventUpdateName = "gorm:update"
)

type (
	callback func(db *gorm.DB)

	Plugin struct {
		createCallback callback
		updateCallback callback
		blackListTable map[string]bool
	}
)

func New() *Plugin {
	plugin := &Plugin{
		createCallback: nil,
		updateCallback: nil,
		blackListTable: map[string]bool{},
	}
	return plugin
}

func (p *Plugin) Name() string {
	return pluginName
}

func (p *Plugin) Initialize(db *gorm.DB) error {
	p.createCallback = p.countryValidation
	p.updateCallback = p.countryValidation

	err := db.
		Callback().
		Create().
		Before(gormEventCreateName).
		Register(customGormEventCreateCallbackName, p.createCallback)
	if err != nil {
		return err
	}

	return db.
		Callback().
		Update().
		After(gormEventUpdateName).
		Register(customGormEventUpdateCallbackName, p.updateCallback)
}

// SetBlacklistTables is function to avoid validation for table
func (p *Plugin) SetBlacklistTables(tablesName []string) {
	for _, tableName := range tablesName {
		if !p.blackListTable[tableName] {
			p.blackListTable[tableName] = true
		}
	}
}

func (p *Plugin) countryValidation(db *gorm.DB) {
	if db.Statement.Schema != nil {
		if !p.blackListTable[db.Statement.Schema.Table] {
			isCountryCodeIsExist := false
			for _, field := range db.Statement.Schema.Fields {
				if field.Name == CountryCode {
					isCountryCodeIsExist = true
					_, isZero := field.ValueOf(context.Background(), db.Statement.ReflectValue)
					if isZero {
						_ = db.AddError(errors.New("value country_code is required please fill it"))
						return
					}
				}
			}

			if !isCountryCodeIsExist {
				_ = db.AddError(errors.New("field country_code is required to declare"))
				return
			}
		}
	}
}
