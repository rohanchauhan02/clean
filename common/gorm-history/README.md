# gorm history
This package is used to keep history of your gorm changes (insert, update, delete). So it acts as a ledger of your gorm model.


## Usage

- Create a history gorm model from the gorm model you want to audit

For example, this is the gorm model that you want to audit
```go
type Product struct {
	ID         int64           `json:"id" db:"id" gorm:"primary_key"`
	Code       string          `json:"code" db:"code" gorm:"column:code"`
	Name       string          `json:"name" db:"name" gorm:"column:name"`
	CategoryID int64           `json:"category_id" db:"category_id" gorm:"column:category_id"`
	Category   *Category       `json:"category,omitempty"`
	Price      float64         `json:"price" db:"price" gorm:"column:price"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time      `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty" db:"deleted_at" gorm:"column:deleted_at"`
}
```

Create the gorm history model of that model
```go
type ProductHistory struct {
	historyLib.Entry
	ID         int64      `gorm:"column:id"`
	Code       string     `gorm:"column:code"`
	Name       string     `gorm:"column:name"`
	CategoryID int64      `gorm:"column:category_id"`
	Price      float64    `gorm:"column:price"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}
```
Your gorm history model should include the `Entry` struct from gorm history library (which will be the fields that tracks the action, user, source, and timestamp details of the history). Also your history model should include the fields that you want to track, with these requirements:

    - Need to have the same name as the original model
    - Need to have the same gorm tag as the original model
    - Need to have the same type as the original model


- Implement implement `Recordable` interface on your gorm model

```go
func (p Product) CreateHistory() historyLib.History {
	return &ProductHistory{}
}
```


- Initiate the gorm history library on your main application

```go
mysql.GetClient().AutoMigrate(models.ProductHiwstory{})
historyPlugin := historyLib.New()
if err := mysql.GetClient().Use(historyPlugin); err != nil {
	msgError := fmt.Sprintf("Failed to initiate gorm history: %s", err.Error())
	e.Logger.Errorf(msgError)
	panic(msgError)
}
```

- Set database history details from context by calling `SetDBHistoryDetails`

```go
db = dbcontext.SetDBHistoryDetails(ctx, db)
```

- After doing these steps, all of the create, update, and delete query will be audited in the history table

## Covered Query Methods

- Create

```go
db.Create(product).Error //single
db.Create(products).Error //multi
```
- Update

```go
db.Model(product).UpdateColumn("price", product.Price).Error //single with the models values thrown
db.Model(&models.Product{}).Where("id = ?", id).UpdateColumn("price", product.Price).Error //single or multiple with only the models thrown without the values
```
- Delete

```go
db.Delete(products).Error //single or multiple with the values thrown
db.Where("id = ?", id).Delete(&models.Product{}).Error // single or multiple with only the models thrown without the values
```

## References

https://github.com/vcraescu/gorm-history
