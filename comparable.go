package check

import (
	"testing"
)

func Equal[Type comparable](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if got != want {
		t.Errorf("%s%T: not equal: got '%v', want '%v'", desc, want, got, want)
		return false
	}

	return true
}
