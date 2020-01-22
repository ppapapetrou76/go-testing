package core

import (
	"fmt"
)

type AnyValue struct {
	value interface{}
}

func (s AnyValue) IsEqualTo(expected interface{}) bool {
	return s.value == expected
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
	case interface{}:
		return AnyValue{value: v}
	case nil:
		return AnyValue{value: v}
	default:
		panic(fmt.Sprintf("expected interface{} value type but got %T type", v))
	}
}
