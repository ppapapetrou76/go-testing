package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/core"
)

// AssertableString is the implementation of CommonAssertable for string types
type AssertableString struct {
	common      core.CommonAssertable
	actual      core.StringValue
}

// ThatString returns an AssertableString structure initialized with the test reference and the actual value to assert
func ThatString(t *testing.T, actual string) AssertableString {
	return AssertableString{
		common: core.CommonAssertable{T: t},
		actual: core.NewStringValue(actual, t),
	}
}

// IsEqualTo asserts if the expected string is equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableString) IsEqualTo(expected string) AssertableString {
	a.common.IsEqualTo(a.actual, expected)
	return a
}

// IsNotEqualTo asserts if the expected string is not equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are equal
func (a AssertableString) IsNotEqualTo(expected string) AssertableString {
	a.common.IsNotEqualTo(a.actual, expected)
	return a
}

// IsGreaterThan asserts if the assertable string value is greater than the expected value
// It errors the tests if is not greater
func (a AssertableString) IsGreaterThan(expected string) AssertableString {
	a.common.IsGreaterThan(a.actual, expected)
	return a
}

// IsGreaterThanOrEqualTo asserts if the assertable int value is greater than or equal to the expected value
// It errors the tests if is not greater
func (a AssertableString) IsGreaterThanOrEqualTo(expected string) AssertableString {
	a.common.IsGreaterOrEqualTo(a.actual, expected)
	return a
}

// IsLessThan asserts if the assertable string value is less than the expected value
// It errors the tests if is not greater
func (a AssertableString) IsLessThan(expected string) AssertableString {
	a.common.IsLessThan(a.actual, expected)
	return a
}

// IsLessThanOrEqualTo asserts if the assertable string value is less than or equal to the expected value
// It errors the tests if is not greater
func (a AssertableString) IsLessThanOrEqualTo(expected string) AssertableString {
	a.common.IsLessOrEqualTo(a.actual, expected)
	return a
}

// IsEmpty asserts if the expected string is empty
// It errors the tests if the string is not empty
func (a AssertableString) IsEmpty() AssertableString {
	a.actual.IsEmpty()
	return a
}

// IsNotEmpty asserts if the expected string is not empty
// It errors the tests if the string is empty
func (a AssertableString) IsNotEmpty() AssertableString {
	a.actual.IsNotEmpty()
	return a
}
