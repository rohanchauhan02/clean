package history

import (
	"reflect"
)

func makePtr(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		return i
	}
	p := reflect.New(v.Type())
	p.Elem().Set(v)
	return p.Interface()
}
