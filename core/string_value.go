package core

import (
	"fmt"
	"strings"
	"testing"
)

type StringValue struct {
	t     *testing.T
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
	return s.greaterThan(NewStringValue(expected, nil))
}

func (s StringValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.greaterOrEqual(NewStringValue(expected, nil))
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

func (s StringValue) Contains(elements ...interface{}) bool {
	for _, e := range elements {
		if !strings.Contains(s.value, NewStringValue(e, nil).value) {
			return false
		}
	}

	return true
}

func (s StringValue) Len() int {
	return len(s.value)
}

func (s StringValue) ContainsOnly(elements ...interface{}) bool {
	return len(s.value) == len(elements) && s.Contains(elements)
}

func (s StringValue) greaterThan(expected StringValue) bool {
	return s.value > expected.value
}

func (s StringValue) greaterOrEqual(expected StringValue) bool {
	return s.value >= expected.value
}

func NewStringValue(value interface{}, t *testing.T) StringValue {
	switch v := value.(type) {
	case string:
		return StringValue{value: v, t: t}
	default:
		panic(fmt.Sprintf("expected string value type but got %T type", v))
	}
}
