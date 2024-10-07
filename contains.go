package check

import (
	"slices"
	"testing"
)

func MapContains[Key comparable, Value any](t *testing.T, m map[Key]Value, want Key, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	_, ok := m[want]
	if !ok {
		t.Errorf("%smap: missing entry '%+v' from '%+v'", desc, want, m)
		return false
	}

	return true
}

func SliceContains[Type comparable](t *testing.T, slice []Type, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !slices.Contains(slice, want) {
		t.Errorf("%sslice: missing entry '%+v' from '%+v'", desc, want, slice)
		return false
	}

	return true
}
