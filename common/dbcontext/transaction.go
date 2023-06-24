package dbcontext

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ctxMysqlClient string

func SetTransactionContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	resDB := db
	ctxValDB := ctx.Value(ctxMysqlClient("mysqlClient"))
	if ctxValDB == nil {
		return resDB
	}
	ctxDB, ok := ctxValDB.(*gorm.DB)
	if ok {
		resDB = ctxDB
	}
	return resDB
}

func GetTransactionContext(ctx context.Context) (*gorm.DB, error) {
	ctxValDB := ctx.Value(ctxMysqlClient("mysqlClient"))
	if ctxValDB == nil {
		return nil, errors.New("mysql client context not found")
	}
	ctxDB, ok := ctxValDB.(*gorm.DB)
	if !ok {
		return nil, errors.New("invalid mysql client context")
	}
	return ctxDB, nil
}

func NewContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxMysqlClient("mysqlClient"), db)
}
