package assert

import (
	"github.com/ppapapetrou76/go-testing/core"
	"testing"
)

// AssertableBool is the assertable structure for bool values
type AssertableBool struct {
	common core.CommonAssertable
	actual core.BoolValue
}

// ThatBool returns an AssertableString structure initialized with the test reference and the actual bool value to assert
func ThatBool(t *testing.T, actual bool) AssertableBool {
	return AssertableBool{
		common: core.CommonAssertable{T: t},
		actual: core.NewBoolValue(actual),
	}
}

// IsEqualTo asserts if the expected bool is equal to the assertable bool value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableBool) IsEqualTo(expected bool) AssertableBool {
	a.common.IsEqualTo(a.actual, expected)
	return a
}

// IsTrue asserts if the expected bool value is true
func (a AssertableBool) IsTrue() AssertableBool {
	return a.IsEqualTo(true)
}

// IsFalse asserts if the expected bool value is false
func (a AssertableBool) IsFalse() AssertableBool {
	return a.IsEqualTo(false)
}
