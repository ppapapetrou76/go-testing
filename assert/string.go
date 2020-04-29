package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

// AssertableString is the implementation of CommonAssertable for string types
type AssertableString struct {
	t      *testing.T
	actual types.StringValue
}

// ThatString returns an AssertableString structure initialized with the test reference and the actual value to assert
func ThatString(t *testing.T, actual string) AssertableString {
	return AssertableString{
		t:      t,
		actual: types.NewStringValue(actual),
	}
}

// IsEqualTo asserts if the expected string is equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableString) IsEqualTo(expected interface{}) AssertableString {
	if a.actual.IsNotEqualTo(expected) {
		a.t.Error(ShouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected string is not equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are equal
func (a AssertableString) IsNotEqualTo(expected interface{}) AssertableString {
	if a.actual.IsEqualTo(expected) {
		a.t.Error(ShouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsEmpty asserts if the expected string is empty
// It errors the tests if the string is not empty
func (a AssertableString) IsEmpty() AssertableString {
	if a.actual.IsNotEmpty() {
		a.t.Error(ShouldBeEmpty(a.actual))
	}
	return a
}

// IsNotEmpty asserts if the expected string is not empty
// It errors the tests if the string is empty
func (a AssertableString) IsNotEmpty() AssertableString {
	if a.actual.IsEmpty() {
		a.t.Error(ShouldNotBeEmpty(a.actual))
	}
	return a
}
