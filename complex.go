package check

import (
	"math"
	"math/cmplx"
	"testing"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func ComplexLess[Type constraints.Complex](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	switch unsafe.Sizeof(want) {
	case 8:
		// 32-bit
		switch {
		// Compare within smallest float32 delta.
		case (cmplx.Abs(complex128(got)) + math.SmallestNonzeroFloat32) >= cmplx.Abs(complex128(want)),
			(cmplx.Abs(complex128(got)) - math.SmallestNonzeroFloat32) >= cmplx.Abs(complex128(want)),
			cmplx.Abs(complex128(got)) >= (cmplx.Abs(complex128(want)) + math.SmallestNonzeroFloat32),
			cmplx.Abs(complex128(got)) >= (cmplx.Abs(complex128(want)) - math.SmallestNonzeroFloat32):
			t.Errorf("%s%T: size not less: got '%v', want less than '%v'", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	case 16:
		// 64-bit
		if !(cmplx.Abs(complex128(got)) < cmplx.Abs(complex128(want))) {
			t.Errorf("%s%T: size not less: got '%v', want less than '%v'", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	default:
		return false
	}

	return true
}

func ComplexLessEqual[Type constraints.Complex](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	switch unsafe.Sizeof(want) {
	case 8:
		// 32-bit
		switch {
		// Compare within smallest float32 delta.
		case (cmplx.Abs(complex128(got)) + math.SmallestNonzeroFloat32) > cmplx.Abs(complex128(want)),
			(cmplx.Abs(complex128(got)) - math.SmallestNonzeroFloat32) > cmplx.Abs(complex128(want)),
			cmplx.Abs(complex128(got)) > (cmplx.Abs(complex128(want)) + math.SmallestNonzeroFloat32),
			cmplx.Abs(complex128(got)) > (cmplx.Abs(complex128(want)) - math.SmallestNonzeroFloat32):
			t.Errorf("%s%T: size not less or equal: got '%v', want '%v' or less", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	case 16:
		// 64-bit
		if !(cmplx.Abs(complex128(got)) <= cmplx.Abs(complex128(want))) {
			t.Errorf("%s%T: size not less or equal: got '%v', want '%v' or less", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	default:
		return false
	}

	return true
}

func ComplexGreater[Type constraints.Complex](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	switch unsafe.Sizeof(want) {
	case 8:
		// 32-bit
		switch {
		// Compare within smallest float32 delta.
		case (cmplx.Abs(complex128(got)) + math.SmallestNonzeroFloat32) <= cmplx.Abs(complex128(want)),
			(cmplx.Abs(complex128(got)) - math.SmallestNonzeroFloat32) <= cmplx.Abs(complex128(want)),
			cmplx.Abs(complex128(got)) <= (cmplx.Abs(complex128(want)) + math.SmallestNonzeroFloat32),
			cmplx.Abs(complex128(got)) <= (cmplx.Abs(complex128(want)) - math.SmallestNonzeroFloat32):
			t.Errorf("%s%T: size not greater: got '%v', want greater than '%v'", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	case 16:
		// 64-bit
		if !(cmplx.Abs(complex128(got)) > cmplx.Abs(complex128(want))) {
			t.Errorf("%s%T: size not greater: got '%v', want greater than '%v'", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	default:
		return false
	}

	return true
}

func ComplexGreaterEqual[Type constraints.Complex](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	switch unsafe.Sizeof(want) {
	case 8:
		// 32-bit
		switch {
		// Compare within smallest float32 delta.
		case (cmplx.Abs(complex128(got)) + math.SmallestNonzeroFloat32) < cmplx.Abs(complex128(want)),
			(cmplx.Abs(complex128(got)) - math.SmallestNonzeroFloat32) < cmplx.Abs(complex128(want)),
			cmplx.Abs(complex128(got)) < (cmplx.Abs(complex128(want)) + math.SmallestNonzeroFloat32),
			cmplx.Abs(complex128(got)) < (cmplx.Abs(complex128(want)) - math.SmallestNonzeroFloat32):
			t.Errorf("%s%T: size not greater or equal: got '%v', want '%v' or greater", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	case 16:
		// 64-bit
		if !(cmplx.Abs(complex128(got)) >= cmplx.Abs(complex128(want))) {
			t.Errorf("%s%T: size not greater or equal: got '%v', want '%v' or greater", desc, want, cmplx.Abs(complex128(got)), cmplx.Abs(complex128(want)))
			return false
		}
	default:
		return false
	}

	return true
}
