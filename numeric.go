package check

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type OrderedNumber interface {
	constraints.Integer | constraints.Float
}

func Delta[Type OrderedNumber](t *testing.T, got, want, delta Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if (got - want) > delta {
		t.Errorf("%s%T: delta exceeded: got '%v', want '%v' or less", desc, delta, got-want, delta)
		return false
	}

	if (want - got) > delta {
		t.Errorf("%s%T: delta exceeded: got '%v', want '%v' or less", desc, delta, want-got, want)
		return false
	}

	return true
}
