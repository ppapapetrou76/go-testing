package values

import (
	"fmt"
	"strings"
)

// StringValue value represents a string value
type StringValue struct {
	value string
}

// IsEqualTo returns true if the value is equal to the expected value, else false
func (s StringValue) IsEqualTo(expected interface{}) bool {
	return s.value == expected
}

// Value returns the actual value of the structure
func (s StringValue) Value() interface{} {
	return s.value
}

// IsGreaterThan returns true if the value is greater than the expected value, else false
func (s StringValue) IsGreaterThan(expected interface{}) bool {
	return s.greaterThan(NewStringValue(expected))
}

// IsGreaterOrEqualTo returns true if the value is greater than or equal to the expected value, else false
func (s StringValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.greaterOrEqual(NewStringValue(expected))
}

// IsLessThan returns true if the value is less than the expected value, else false
func (s StringValue) IsLessThan(expected interface{}) bool {
	return !s.IsGreaterOrEqualTo(expected)
}

// IsLessOrEqualTo returns true if the value is less than or equal to the expected value, else false
func (s StringValue) IsLessOrEqualTo(expected interface{}) bool {
	return !s.IsGreaterThan(expected)
}

// IsEmpty returns true if the string is empty else false
func (s StringValue) IsEmpty() bool {
	return s.value == ""
}

// IsNotEmpty returns true if the string is not empty else false
func (s StringValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Contains returns true if the string contains the given sub-string
func (s StringValue) Contains(element interface{}) bool {
	return strings.Contains(s.value, NewStringValue(element).value)
}

// DoesNotContain returns true if the string does not contain the given sub-string
func (s StringValue) DoesNotContain(elements interface{}) bool {
	return !s.Contains(elements)
}

// HasSize returns true if the string has the expected size else false
func (s StringValue) HasSize(length int) bool {
	return s.Size() == length
}

// Size returns the string size
func (s StringValue) Size() int {
	return len(s.value)
}

// ContainsOnly returns true if the string contains only the given sub-string
// In other words if performs an equal operation
func (s StringValue) ContainsOnly(elements interface{}) bool {
	return s.IsEqualTo(elements)
}

func (s StringValue) greaterThan(expected StringValue) bool {
	return s.value > expected.value
}

func (s StringValue) greaterOrEqual(expected StringValue) bool {
	return s.value >= expected.value
}

// NewStringValue creates and returns a StringValue struct initialed with the given value
func NewStringValue(value interface{}) StringValue {
	switch v := value.(type) {
	case string:
		return StringValue{value: v}
	default:
		panic(fmt.Sprintf("expected string value type but got %T type", v))
	}
}
