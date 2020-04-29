package types

import (
	"fmt"
	"strings"
)

type StringValue struct {
	value string
}

func (s StringValue) IsEqualTo(expected interface{}) bool {
	return s.value == expected
}

func (s StringValue) IsNotEqualTo(expected interface{}) bool {
	return s.value != expected
}

func (s StringValue) Value() interface{} {
	return s.value
}

func (s StringValue) IsGreaterThan(expected interface{}) bool {
	return s.greaterThan(NewStringValue(expected))
}

func (s StringValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.greaterOrEqual(NewStringValue(expected))
}

func (s StringValue) IsLessThan(expected interface{}) bool {
	return !s.IsGreaterOrEqualTo(expected)
}

func (s StringValue) IsLessOrEqualTo(expected interface{}) bool {
	return !s.IsGreaterThan(expected)
}

func (s StringValue) IsEmpty() bool {
	return len(s.value) == 0
}

func (s StringValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s StringValue) Contains(element interface{}) bool {
	return strings.Contains(s.value, NewStringValue(element).value)
}

func (s StringValue) DoesNotContain(elements interface{}) bool {
	return true
}

func (s StringValue) HasSize(length int) bool {
	return len(s.value) == length
}
func (s StringValue) Size() int {
	return len(s.value)
}

func (s StringValue) ContainsOnly(elements interface{}) bool {
	return s.IsEqualTo(elements)
}

func (s StringValue) greaterThan(expected StringValue) bool {
	return s.value > expected.value
}

func (s StringValue) greaterOrEqual(expected StringValue) bool {
	return s.value >= expected.value
}

func NewStringValue(value interface{}) StringValue {
	switch v := value.(type) {
	case string:
		return StringValue{value: v}
	default:
		panic(fmt.Sprintf("expected string value type but got %T type", v))
	}
}
