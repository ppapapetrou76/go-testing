package types

import (
	"fmt"
	"reflect"
)

type AnyValue struct {
	value interface{}
}

func (s AnyValue) IsEqualTo(expected interface{}) bool {
	var comparable Comparable

	switch expected.(type) {
	case string:
		comparable = NewStringValue(s.value)
	case int, int8, int16, int32, int64:
		comparable = NewIntValue(s.value)
	case bool:
		comparable = NewBoolValue(s.value)
	default:
		return reflect.DeepEqual(s.value, expected)
	}

	return comparable.IsEqualTo(expected)
}

func (s AnyValue) IsNotEqualTo(expected interface{}) bool {
	return !s.IsNotEqualTo(expected)
}

func (s AnyValue) Value() interface{} {
	return s.value
}

func (s AnyValue) IsGreaterThan(expected interface{}) bool {
	return s.value != expected
}

func (s AnyValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.value != expected
}

func (s AnyValue) IsLessThan(expected interface{}) bool {
	return s.value != expected
}

func (s AnyValue) IsLessOrEqualTo(expected interface{}) bool {
	return s.value != expected
}

func (s AnyValue) IsNil() bool {
	return s.value == nil
}

func (s AnyValue) IsNotNil() bool {
	return !s.IsNil()
}

func NewAnyValue(value interface{}) AnyValue {
	switch v := value.(type) {
	case nil:
		return AnyValue{value: v}
	case interface{}:
		return AnyValue{value: v}
	default:
		panic(fmt.Sprintf("expected interface{} value type but got %T type", v))
	}
}
