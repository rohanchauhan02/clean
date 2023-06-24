package history

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	ActionCreate    Action             = "create"
	ActionUpdate    Action             = "update"
	ActionDelete    Action             = "delete"
	userOptionKey   userOptionCtxKey   = pluginName + ":user"
	sourceOptionKey sourceOptionCtxKey = pluginName + ":source"
)

var (
	_ History              = (*Entry)(nil)
	_ TimestampableHistory = (*Entry)(nil)
	_ BlameableHistory     = (*Entry)(nil)
	_ SourceableHistory    = (*Entry)(nil)
)

type (
	Action string

	userOptionCtxKey string

	sourceOptionCtxKey string

	Recordable interface {
		CreateHistory() History
	}

	BlameableHistory interface {
		SetHistoryUserID(id string)
		SetHistoryUserType(userType string)
		SetHistoryUserEmail(email string)
		SetHistoryUserFullName(fullName string)
	}

	SourceableHistory interface {
		SetHistorySourceRequestID(requestID string)
	}

	TimestampableHistory interface {
		SetHistoryCreatedAt(createdAt time.Time)
	}

	History interface {
		SetHistoryAction(action Action)
	}

	Entry struct {
		DBHistoryID           int64     `gorm:"column:db_history_id;primaryKey"`
		DBHistoryAction       Action    `gorm:"column:db_history_action;type:varchar(24)"`
		DBHistoryUserID       string    `gorm:"column:db_history_user_id;type:varchar(255)"`
		DBHistoryUserType     string    `gorm:"column:db_history_user_type;type:varchar(255)"`
		DBHistoryUserEmail    string    `gorm:"column:db_history_user_email;type:varchar(255)"`
		DBHistoryUserFullName string    `gorm:"column:db_history_user_full_name;type:varchar(255)"`
		DBHistoryRequestID    string    `gorm:"column:db_history_request_id;type:varchar(255)"`
		DBHistoryCreatedAt    time.Time `gorm:"column:db_history_created_at;type:datetime"`
	}

	User struct {
		ID       string
		Type     string
		FullName string
		Email    string
	}

	Source struct {
		RequestID string
	}
)

func SetUser(db *gorm.DB, user User) *gorm.DB {
	ctx := context.WithValue(db.Statement.Context, userOptionKey, user)
	return db.WithContext(ctx).Set(string(userOptionKey), user)
}

func GetUser(db *gorm.DB) (User, bool) {
	value, ok := db.Get(string(userOptionKey))
	if !ok {
		value := db.Statement.Context.Value(userOptionKey)
		user, ok := value.(User)

		return user, ok
	}
	user, ok := value.(User)
	return user, ok
}

func SetSource(db *gorm.DB, source Source) *gorm.DB {
	ctx := context.WithValue(db.Statement.Context, sourceOptionKey, source)
	return db.WithContext(ctx).Set(string(sourceOptionKey), source)
}

func GetSource(db *gorm.DB) (Source, bool) {
	value, ok := db.Get(string(sourceOptionKey))
	if !ok {
		value := db.Statement.Context.Value(sourceOptionKey)
		source, ok := value.(Source)

		return source, ok
	}
	source, ok := value.(Source)
	return source, ok
}

func (e *Entry) SetHistoryAction(action Action) {
	e.DBHistoryAction = action
}

func (e *Entry) SetHistoryUserID(id string) {
	e.DBHistoryUserID = id
}

func (e *Entry) SetHistoryUserType(userType string) {
	e.DBHistoryUserType = userType
}

func (e *Entry) SetHistoryUserEmail(email string) {
	e.DBHistoryUserEmail = email
}

func (e *Entry) SetHistoryUserFullName(fullName string) {
	e.DBHistoryUserFullName = fullName
}

func (e *Entry) SetHistorySourceRequestID(id string) {
	e.DBHistoryRequestID = id
}

func (e *Entry) SetHistoryCreatedAt(createdAt time.Time) {
	e.DBHistoryCreatedAt = createdAt
}
