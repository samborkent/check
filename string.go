package check

import (
	"strings"
	"testing"
)

func StringEqualFold(t *testing.T, got, want string, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !strings.EqualFold(got, want) {
		t.Errorf("%sstring: not equal under case-fold: got '%s', want '%s'", desc, got, want)
		return false
	}

	return true
}
