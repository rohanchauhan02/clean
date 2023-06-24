package dbcontext

import (
	"context"

	"gorm.io/gorm"
)

func SetTracer(ctx context.Context, db *gorm.DB) *gorm.DB {
	if ctx == nil {
		return db
	}
	db = db.WithContext(ctx)
	return db
}
