package check

import (
	"errors"
	"testing"
)

func ErrorEqual(t *testing.T, got, want error, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !errors.Is(got, want) {
		t.Errorf("%serror: not equal: got '%s', want '%s'", desc, got.Error(), want.Error())
		return false
	}

	return true
}

func ErrorNil(t *testing.T, err error, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if err != nil {
		t.Errorf("%serror: not nil: got '%s'", desc, err.Error())
		return false
	}

	return true
}
