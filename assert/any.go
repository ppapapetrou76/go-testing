package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

// Assertable is the assertable structure for interface{} values
type Assertable struct {
	t      *testing.T
	actual types.AnyValue
}

// That returns an Assertable structure initialized with the test reference and the actual value to assert
func That(t *testing.T, actual interface{}) Assertable {
	return Assertable{
		t:      t,
		actual: types.NewAnyValue(actual),
	}
}

// IsEqualTo asserts if the expected interface is equal to the assertable value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a Assertable) IsEqualTo(expected interface{}) Assertable {
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected interface is not qual to the assertable value
// It errors the tests if the compared values (actual VS expected) are equal
func (a Assertable) IsNotEqualTo(expected interface{}) Assertable {
	if !a.actual.IsNotEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsNil asserts if the expected value is nil
func (a Assertable) IsNil() Assertable {
	if !a.actual.IsNil() {
		a.t.Error(shouldBeNil(a.actual))
	}
	return a
}

// IsNotNil asserts if the expected value is not nil
func (a Assertable) IsNotNil() Assertable {
	if !a.actual.IsNotNil() {
		a.t.Error(shouldNotBeNil(a.actual))
	}
	return a
}

// IsTrue asserts if the expected value is true
func (a Assertable) IsTrue() Assertable {
	a.IsEqualTo(true)
	return a
}

// IsFalse asserts if the expected value is false
func (a Assertable) IsFalse() Assertable {
	a.IsEqualTo(false)
	return a
}
