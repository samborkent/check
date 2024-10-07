package check

import (
	"reflect"
	"testing"
	"unsafe"
)

func AnyEqual[Type any](t *testing.T, got, want Type, descriptions ...string) bool {
	t.Helper()

	desc := joinDescriptions(descriptions...)

	if unsafe.Sizeof(got) != unsafe.Sizeof(want) {
		t.Errorf("%ssize mismatch: got '%d', want '%d'", desc, unsafe.Sizeof(got), unsafe.Sizeof(want))
		return false
	}

	gotValue := reflect.ValueOf(got)
	wantValue := reflect.ValueOf(want)

	gotValid := gotValue.IsValid()
	wantValid := wantValue.IsValid()

	if gotValid != wantValid {
		t.Errorf("%svalidity mismatch: got '%t', want '%t'", desc, gotValid, wantValid)
		return false
	}

	gotType := gotValue.Type()
	wantType := wantValue.Type()

	if gotType != wantType {
		t.Errorf("%stype mismatch: got '%s', want '%s'", desc, gotType.String(), wantType.String())
		return false
	}

	gotKind := gotValue.Kind()
	wantKind := gotValue.Kind()

	if gotKind != wantKind {
		t.Errorf("%stype kind mismatch: got '%s', want '%s'", desc, gotKind.String(), wantKind.String())
		return false
	}

	switch gotKind {
	case reflect.Array:
		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%sarray len mismatch: got '%d', want '%d'", desc, gotValue.Len(), wantValue.Len())
			return false
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%sarray mismatch: got '%v', want '%v'", desc, got, want)
			return false
		}
	case reflect.Bool:
		return Equal(t, gotValue.Bool(), wantValue.Bool())
	case reflect.Chan:
		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%schan len mismatch: got '%d', want '%d'", desc, gotValue.Len(), wantValue.Len())
			return false
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%schan mismatch: got '%v', want '%v'", desc, gotValue, wantValue)
			return false
		}
	case reflect.Complex64, reflect.Complex128:
		return Equal(t, gotValue.Complex(), wantValue.Complex())
	case reflect.Float32, reflect.Float64:
		return Equal(t, gotValue.Float(), wantValue.Float())
	case reflect.Func:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%sfunc mismatch: got '%+v', want '%+v'", desc, gotValue, wantValue)
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Equal(t, gotValue.Int(), wantValue.Int())
	case reflect.Interface:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%sinterface mismatch: got '%+v', want '%+v'", desc, gotValue, wantValue)
			return false
		}
	case reflect.Map:
		if gotValue.IsNil() != wantValue.IsNil() {
			t.Errorf("%snil map mismatch: got '%t', want '%t'", desc, gotValue.IsNil(), wantValue.IsNil())
			return false
		}

		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%smap len mismatch: got '%d', want '%d'", desc, gotValue.Len(), wantValue.Len())
			return false
		}

		gotPtr := gotValue.UnsafePointer()
		wantPtr := wantValue.UnsafePointer()

		if gotPtr != wantPtr {
			t.Errorf("%smap pointer mismatch: got '%v', want '%v'", desc, gotPtr, wantPtr)
			return false
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%smap mismatch: got '%v', want '%v'", desc, got, want)
			return false
		}
	case reflect.Pointer:
		gotPtr := gotValue.UnsafePointer()
		wantPtr := wantValue.UnsafePointer()

		if gotPtr != wantPtr {
			t.Errorf("%sunsafe pointer mismatch: got '%v', want '%v'", desc, gotPtr, wantPtr)
			return false
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%spointer mismatch: got '%+v', want '%+v'", desc, gotValue, wantValue)
			return false
		}
	case reflect.Slice:
		if gotValue.IsNil() != wantValue.IsNil() {
			t.Errorf("%snil slice mismatch: got '%t', want '%t'", desc, gotValue.IsNil(), wantValue.IsNil())
			return false
		}

		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%smap len mismatch: got '%d', want '%d'", desc, gotValue.Len(), wantValue.Len())
			return false
		}

		gotPtr := gotValue.UnsafePointer()
		wantPtr := wantValue.UnsafePointer()

		if gotPtr != wantPtr {
			t.Errorf("%smap pointer mismatch: got '%v', want '%v'", desc, gotPtr, wantPtr)
			return false
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%sslice mismatch: got '%v', want '%v'", desc, got, want)
			return false
		}
	case reflect.Struct:
		if gotValue.NumField() != wantValue.NumField() {
			t.Errorf("%sstruct field number mismatch: got '%+v', want '%+v'", desc, gotValue.NumField(), wantValue.NumField())
			return false
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%sstruct mismatch: got '%+v', want '%+v'", desc, got, want)
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Equal(t, gotValue.Uint(), wantValue.Uint())
	default:
		gotAny := gotValue.Interface()
		wantAny := wantValue.Interface()

		if gotAny != wantAny {
			t.Errorf("%suncasted interface mismatch: got '%+v', want '%+v'", desc, gotValue, wantValue)
			return false
		}
	}

	return true
}
