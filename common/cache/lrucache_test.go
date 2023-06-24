package cache

import "testing"

func TestAddCache(t *testing.T) {
	l := New(10)

	type testStruct struct {
		ID   string
		Name string
	}
	data := []testStruct{
		testStruct{
			ID:   "ID-1",
			Name: "Name 1",
		},
		testStruct{
			ID:   "ID-2",
			Name: "Name 2",
		},
	}
	key := "TEMP_KEY"
	l.Set(key, data)

	valRes, ok := l.Get(key)
	if ok {
		t.Log(valRes.([]testStruct))
		return
	}
	t.Error("Should not error")
}
