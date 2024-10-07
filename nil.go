package check

import (
	"reflect"
	"testing"
)

func Nil[Type any](t *testing.T, got Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if any(got) != nil || !reflect.ValueOf(got).IsNil() {
		t.Errorf("%s%T: not nil: got '%+v'", desc, got, got)
		return false
	}

	return true
}
