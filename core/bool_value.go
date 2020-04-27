package core

import (
	"fmt"
)

type BoolValue struct {
	value bool
}

func (s BoolValue) IsEqualTo(expected interface{}) bool {
	return s.value == expected
}

func (s BoolValue) IsNotEqualTo(expected interface{}) bool {
	return s.value != expected
}

func (s BoolValue) Value() interface{} {
	return s.value
}

func (s BoolValue) IsGreaterThan(expected interface{}) bool {
	return s.value != expected
}

func (s BoolValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.value != expected
}

func (s BoolValue) IsLessThan(expected interface{}) bool {
	return s.value != expected
}

func (s BoolValue) IsLessOrEqualTo(expected interface{}) bool {
	return s.value != expected
}

func NewBoolValue(value interface{}) BoolValue {
	switch v := value.(type) {
	case bool:
		return BoolValue{value: v}
	default:
		panic(fmt.Sprintf("expected bool value type but got %T type", v))
	}
}
