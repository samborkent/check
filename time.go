package check

import (
	"testing"
	"time"
)

func TimeAfter(t *testing.T, got, want time.Time, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !got.After(want) {
		t.Errorf("%stime.Time: not after: got '%s', want '%s'", desc, got.Format(time.RFC3339Nano), want.Format(time.RFC3339Nano))
		return false
	}

	return true
}

func TimeBefore(t *testing.T, got, want time.Time, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !got.Before(want) {
		t.Errorf("%stime.Time: not before: got '%s', want '%s'", desc, got.Format(time.RFC3339Nano), want.Format(time.RFC3339Nano))
		return false
	}

	return true
}

func TimeDelta(t *testing.T, got, want time.Time, delta time.Duration, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if got.Sub(want).Abs() > delta {
		t.Errorf("%stime.Time: delta exceeded: got '%d', want '%s' or less", desc, got.Sub(want).Abs(), delta)
		return false
	}

	return true
}

func TimeEqual(t *testing.T, got, want time.Time, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !got.Equal(want) {
		t.Errorf("%stime.Time: not equal: got '%s', want '%s'", desc, got.Format(time.RFC3339Nano), want.Format(time.RFC3339Nano))
		return false
	}

	return true
}

func TimeZero(t *testing.T, got time.Time, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if !got.IsZero() {
		t.Errorf("%stime.Time: not zero: got '%s', want '%s'", desc, got.Format(time.RFC3339Nano), time.Time{}.Format(time.RFC3339Nano))
		return false
	}

	return true
}
