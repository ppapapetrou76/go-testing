package assert

import (
	"github.com/ppapapetrou76/go-testing/core"
	"testing"
)

// Assertable is the assertable structure for interface{} values
type Assertable struct {
	common core.CommonAssertable
	actual core.AnyValue
}

// That returns an Assertable structure initialized with the test reference and the actual value to assert
func That(t *testing.T, actual interface{}) Assertable {
	return Assertable{
		common: core.CommonAssertable{T: t},
		actual: core.NewAnyValue(actual),
	}
}

// IsEqualTo asserts if the expected interface is equal to the assertable value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a Assertable) IsEqualTo(expected bool) Assertable {
	a.common.IsEqualTo(a.actual, expected)
	return a
}

// IsNil asserts if the expected value is nil
func (a Assertable) IsNil() Assertable {
	a.common.IsNil(a.actual)
	return a
}

// IsNotNil asserts if the expected value is not nil
func (a Assertable) IsNotNil() Assertable {
	a.common.IsNotNil(a.actual)
	return a
}
