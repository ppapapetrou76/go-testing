package assert

import (
	"github.com/ppapapetrou76/go-testing/core"
	"testing"
)

// AssertableInt
type AssertableInt struct {
	common core.CommonAssertable
	actual core.IntValue
}

// ThatInt returns an AssertableInt structure initialized with the test reference and the actual value to assert
func ThatInt(t *testing.T, actual int) AssertableInt {
	return AssertableInt{
		common: core.CommonAssertable{T: t},
		actual: core.NewIntValue(actual),
	}
}

// IsEqualTo asserts if the expected int is equal to the assertable int value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableInt) IsEqualTo(expected int) AssertableInt {
	a.common.IsEqualTo(a.actual, expected)
	return a
}

// IsNotEqualTo asserts if the expected int is not equal to the assertable int value
// It errors the tests if the compared values (actual VS expected) are equal
func (a AssertableInt) IsNotEqualTo(expected int) AssertableInt {
	a.common.IsNotEqualTo(a.actual, expected)
	return a
}

// IsGreaterThan asserts if the assertable int value is greater than the expected value
// It errors the tests if is not greater
func (a AssertableInt) IsGreaterThan(expected int) AssertableInt {
	a.common.IsGreaterThan(a.actual, expected)
	return a
}

// IsGreaterThanOrEqualTo asserts if the assertable int value is greater than or equal to the expected value
// It errors the tests if is not greater
func (a AssertableInt) IsGreaterThanOrEqualTo(expected int) AssertableInt {
	a.common.IsGreaterOrEqualTo(a.actual, expected)
	return a
}

// IsLessThan asserts if the assertable int value is less than the expected value
// It errors the tests if is not greater
func (a AssertableInt) IsLessThan(expected int) AssertableInt {
	a.common.IsLessThan(a.actual, expected)
	return a
}

// IsLessThanOrEqualTo asserts if the assertable int value is less than or equal to the expected value
// It errors the tests if is not greater
func (a AssertableInt) IsLessThanOrEqualTo(expected int) AssertableInt {
	a.common.IsLessOrEqualTo(a.actual, expected)
	return a
}
