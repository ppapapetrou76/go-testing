package core

import (
	"fmt"
)

type IntValue struct {
	value int
}

func (i IntValue) IsEqualTo(expected interface{}) bool {
	return i.value == expected
}

func (i IntValue) IsNotEqualTo(expected interface{}) bool {
	return i.value != expected
}

func (i IntValue) IsGreaterThan(expected interface{}) bool {
	return i.greaterThan(NewIntValue(expected))
}

func (i IntValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return i.greaterOrEqual(NewIntValue(expected))
}

func (i IntValue) IsLessThan(expected interface{}) bool {
	return !i.IsGreaterOrEqualTo(expected)
}

func (i IntValue) IsLessOrEqualTo(expected interface{}) bool {
	return !i.IsGreaterThan(expected)
}

func (i IntValue) Value() interface{} {
	return i.value
}

func (i IntValue) greaterThan(expected IntValue) bool {
	return i.value > expected.value
}

func (i IntValue) greaterOrEqual(expected IntValue) bool {
	return i.value >= expected.value
}

func NewIntValue(value interface{}) IntValue {
	switch v := value.(type) {
	case int:
		return IntValue{value: v}
	default:
		panic(fmt.Sprintf("expected int value type but got %T type", v))
	}
}
