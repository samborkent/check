package check

import (
	"reflect"
	"testing"
)

func Nil[Type any](t *testing.T, got Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if any(got) == nil {
		return true
	}

	value := reflect.ValueOf(got)

	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if value.IsNil() {
			return true
		}
	}

	t.Errorf("%s%s: not nil: got '%+v'", value.Type().String(), desc, got)
	return false
}
