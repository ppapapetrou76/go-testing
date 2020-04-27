package core

import (
	"fmt"
	"reflect"
)

type AnyValue struct {
	value interface{}
}

func (s AnyValue) IsEqualTo(expected interface{}) bool {
	switch expected.(type) {
	case string:
		return s.value == expected
	case int:
		return s.value == expected
	case int8:
		return s.value == expected
	case int16:
		return s.value == expected
	case int32:
		return s.value == expected
	case int64:
		return s.value == expected
	case bool:
		return s.value == expected
	default:
		return reflect.DeepEqual(s.value, expected)
	}
}

func (s AnyValue) IsNotEqualTo(expected interface{}) bool {
	return s.value != expected
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
