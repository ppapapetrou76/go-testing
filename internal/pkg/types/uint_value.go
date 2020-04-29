package types

import (
	"fmt"
)

type UIntValue struct {
	value uint
}

func (i UIntValue) IsEqualTo(expected interface{}) bool {
	return i.value == expected
}

func (i UIntValue) IsNotEqualTo(expected interface{}) bool {
	return i.value != expected
}

func (i UIntValue) IsGreaterThan(expected interface{}) bool {
	return i.greaterThan(NewUIntValue(expected))
}

func (i UIntValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return i.greaterOrEqual(NewUIntValue(expected))
}

func (i UIntValue) IsLessThan(expected interface{}) bool {
	return !i.IsGreaterOrEqualTo(expected)
}

func (i UIntValue) IsLessOrEqualTo(expected interface{}) bool {
	return !i.IsGreaterThan(expected)
}

func (i UIntValue) Value() interface{} {
	return i.value
}

func (i UIntValue) greaterThan(expected UIntValue) bool {
	return i.value > expected.value
}

func (i UIntValue) greaterOrEqual(expected UIntValue) bool {
	return i.value >= expected.value
}

func NewUIntValue(value interface{}) UIntValue {
	switch v := value.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return UIntValue{value: value.(uint)}
	default:
		panic(fmt.Sprintf("expected uint value type but got %T type", v))
	}
}
