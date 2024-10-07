package check

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func Less[Type constraints.Ordered](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !(got < want) {
		t.Errorf("%s%T: not less: got '%v', want less than '%v'", desc, want, got, want)
		return false
	}

	return true
}

func LessEqual[Type constraints.Ordered](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !(got <= want) {
		t.Errorf("%s%T: not less or equal: got '%v', want '%v' or less than", desc, want, got, want)
		return false
	}

	return true
}

func Greater[Type constraints.Ordered](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !(got > want) {
		t.Errorf("%s%T: not greater: got '%v', want greater than '%v'", desc, want, got, want)
		return false
	}

	return true
}

func GreaterEqual[Type constraints.Ordered](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !(got >= want) {
		t.Errorf("%s%T: not greater or equal: got '%v', want '%v' or greater than", desc, want, got, want)
		return false
	}

	return true
}
