package history

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/jinzhu/copier"
)

const (
	pluginName                             = "gorm-history"
	createCbName                           = pluginName + ":after_create"
	updateCbName                           = pluginName + ":after_update"
	deleteCbName                           = pluginName + ":before_delete"
	disabledOptionKey disabledOptionCtxKey = pluginName + ":disabled"
)

var (
	_ gorm.Plugin = (*Plugin)(nil)

	ErrUnsupportedOperation = errors.New("history is not supported for this operation")
)

type (
	disabledOptionCtxKey string

	CopyFunc func(r Recordable, h interface{}) error

	callback func(db *gorm.DB)

	Context struct {
		object  Recordable
		history History
		action  Action
		db      *gorm.DB
	}

	Config struct {
		CopyFunc CopyFunc
	}

	ConfigFunc func(c *Config)

	Plugin struct {
		copyFunc CopyFunc
		createCb callback
		updateCb callback
		deleteCb callback
	}
)

func New(configFuncs ...ConfigFunc) *Plugin {
	cfg := &Config{
		CopyFunc: DefaultCopyFunc,
	}

	for _, f := range configFuncs {
		f(cfg)
	}

	p := Plugin{
		copyFunc: cfg.CopyFunc,
	}

	return &p
}

func WithCopyFunc(fn CopyFunc) ConfigFunc {
	return func(c *Config) {
		c.CopyFunc = fn
	}
}

func Disable(db *gorm.DB) *gorm.DB {
	ctx := context.WithValue(db.Statement.Context, disabledOptionKey, true)
	return db.WithContext(ctx).Set(string(disabledOptionKey), true)
}

func IsDisabled(db *gorm.DB) bool {
	_, ok := db.Get(string(disabledOptionKey))
	if ok {
		return true
	}
	val := db.Statement.Context.Value(disabledOptionKey)
	return val != nil
}

func (p *Plugin) Name() string {
	return pluginName
}

func (p *Plugin) Initialize(db *gorm.DB) error {
	p.createCb = p.callback(ActionCreate)
	p.updateCb = p.callback(ActionUpdate)
	p.deleteCb = p.callback(ActionDelete)

	err := db.
		Callback().
		Create().
		After("gorm:create").
		Register(createCbName, p.createCb)
	if err != nil {
		return err
	}

	err = db.
		Callback().
		Update().
		After("gorm:update").
		Register(updateCbName, p.updateCb)
	if err != nil {
		return err
	}

	return db.
		Callback().
		Delete().
		Before("gorm:delete").
		Register(deleteCbName, p.deleteCb)
}

func (p Plugin) callback(action Action) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if db.Statement.Schema == nil {
			return
		}

		if IsDisabled(db) {
			return
		}

		v := db.Statement.ReflectValue

		if action == ActionDelete || action == ActionUpdate {
			h, isRecordable, err := p.processStructUpdateDelete(v, action, db)
			if err != nil {
				db.AddError(err)
				return
			}

			if !isRecordable {
				return
			}
			if err := p.saveHistory(db, h...); err != nil {
				db.AddError(err)
				return
			}
			return
		}

		// for create action only
		switch v.Kind() {
		case reflect.Struct:
			h, isRecordable, err := p.processStruct(v, action, db)
			if err != nil {
				db.AddError(err)
				return
			}

			if !isRecordable {
				return
			}
			if err := p.saveHistory(db, h); err != nil {
				db.AddError(err)
				return
			}
		case reflect.Slice:
			hs, err := p.processSlice(v, action, db)
			if err != nil {
				db.AddError(err)
				return
			}

			if len(hs) == 0 {
				return
			}

			if err := p.saveHistory(db, hs...); err != nil {
				db.AddError(err)
				return
			}
		}

	}
}

func (p *Plugin) saveHistory(db *gorm.DB, hs ...History) error {
	if len(hs) == 0 {
		return nil
	}
	db = db.Session(&gorm.Session{
		NewDB: true,
	})
	for _, h := range hs {
		if err := db.Omit(clause.Associations).Create(h).Error; err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) processStructUpdateDelete(v reflect.Value, action Action, db *gorm.DB) ([]History, bool, error) {
	var err error
	vi := v.Interface()
	var arrModel interface{}

	if v.Kind() == reflect.Slice {
		arrModel = reflect.New(reflect.SliceOf(v.Type().Elem())).Interface()
	} else {
		arrModel = reflect.New(reflect.SliceOf(reflect.TypeOf(vi))).Interface()
	}

	// prepare where clause
	var updateClause string
	var updateVars []interface{}
	queryStmt := db.Session(&gorm.Session{DryRun: true}).Statement
	updateQuery := queryStmt.SQL.String()
	splitClauses := strings.Split(updateQuery, "WHERE ")
	if len(splitClauses) == 2 {
		updateClause = splitClauses[1]
		whereArgsIdx := strings.Count(strings.Split(updateQuery, "WHERE ")[0], "?")
		updateVars = queryStmt.Vars[whereArgsIdx:]
	} else {
		err = errors.New("invalid where condition provided")
		return nil, false, err
	}

	db = db.Session(&gorm.Session{
		NewDB: true,
	})
	db = db.Where(updateClause, updateVars)
	err = db.Find(arrModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}

	var histories []History
	arrModelElm := reflect.ValueOf(arrModel).Elem()
	for i := 0; i < arrModelElm.Len(); i++ {
		objReflect := arrModelElm.Index(i)
		objPointerReflect := reflect.NewAt(objReflect.Type(), unsafe.Pointer(objReflect.UnsafeAddr()))
		objInterface := objPointerReflect.Interface()
		r, ok := objInterface.(Recordable)
		if !ok {
			return nil, false, nil
		}
		h, err := p.newHistory(r, action, db)
		if err != nil {
			return nil, false, err
		}
		histories = append(histories, h)
	}

	return histories, true, nil
}

func (p *Plugin) processStruct(v reflect.Value, action Action, db *gorm.DB) (History, bool, error) {
	vi := v.Interface()
	r, ok := vi.(Recordable)
	if !ok {
		return nil, false, nil
	}

	var err error
	h, err := p.newHistory(r, action, db)
	if err != nil {
		return nil, true, err
	}

	return h, true, nil
}

func (p *Plugin) processSlice(v reflect.Value, action Action, db *gorm.DB) ([]History, error) {
	var hs []History
	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)

		h, isRecordable, err := p.processStruct(el, action, db)
		if err != nil {
			return nil, err
		}

		if !isRecordable {
			continue
		}

		hs = append(hs, h)
	}

	return hs, nil
}

func (p *Plugin) newHistory(r Recordable, action Action, db *gorm.DB) (History, error) {
	hist := r.CreateHistory()
	ihist := makePtr(hist)
	if err := p.copyFunc(r, ihist); err != nil {
		return nil, err
	}

	if err := db.Statement.Parse(hist); err != nil {
		return nil, err
	}

	hist.SetHistoryAction(action)

	if th, ok := hist.(TimestampableHistory); ok {
		th.SetHistoryCreatedAt(db.NowFunc())
	}

	if bh, ok := hist.(BlameableHistory); ok {
		if user, ok := GetUser(db); ok {
			bh.SetHistoryUserID(user.ID)
			bh.SetHistoryUserType(user.Type)
			bh.SetHistoryUserEmail(user.Email)
			bh.SetHistoryUserFullName(user.FullName)
		}
	}

	if bh, ok := hist.(SourceableHistory); ok {
		if source, ok := GetSource(db); ok {
			bh.SetHistorySourceRequestID(source.RequestID)
		}
	}

	return hist.(History), nil
}

func DefaultCopyFunc(r Recordable, h interface{}) error {
	if reflect.ValueOf(h).Kind() != reflect.Ptr {
		return fmt.Errorf("pointer expected but got %T", h)
	}
	return copier.Copy(h, r)
}
