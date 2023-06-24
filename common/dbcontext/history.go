package dbcontext

import (
	"context"

	historyLib "github.com/rohanchauhan02/clean/common/gorm-history"
	middlewareLib "github.com/rohanchauhan02/clean/common/middleware"
	"gorm.io/gorm"
)

func SetDBHistoryDetails(ctx context.Context, db *gorm.DB) *gorm.DB {
	var historyUserData historyLib.User
	var historySourceData historyLib.Source

	if ctx == nil {
		return db
	}

	ctxValDB := ctx.Value(middlewareLib.ContextHistoryUserKey)
	if ctxValDB != nil {
		var ok bool
		historyUserData, ok = ctxValDB.(historyLib.User)
		if !ok {
			return db
		}
	}

	ctxValDB = ctx.Value(middlewareLib.ContextHistorySourceKey)
	if ctxValDB != nil {
		var ok bool
		historySourceData, ok = ctxValDB.(historyLib.Source)
		if !ok {
			return db
		}
	}

	db = historyLib.SetUser(db, historyLib.User{
		ID:       historyUserData.ID,
		Type:     historyUserData.Type,
		FullName: historyUserData.FullName,
		Email:    historyUserData.Email,
	})

	db = historyLib.SetSource(db, historyLib.Source{
		RequestID: historySourceData.RequestID,
	})

	return db
}
