package values

import (
	"fmt"
	"reflect"

	"github.com/ppapapetrou76/go-testing/types"
)

// AnyValue is a struct that holds any type of value
type AnyValue struct {
	value interface{}
}

// IsEqualTo returns true if the value is equal to the expected value, else false
func (s AnyValue) IsEqualTo(expected interface{}) bool {
	var comparable types.Comparable

	switch expected.(type) {
	case string:
		comparable = NewStringValue(s.value)
	case int, int8, int16, int32, int64:
		comparable = NewIntValue(s.value)
	case uint, uint8, uint16, uint32, uint64:
		comparable = NewUIntValue(s.value)
	case bool:
		comparable = NewBoolValue(s.value)
	case struct{}:
		comparable = NewStructValue(s.value)
	case []string:
		comparable = NewSliceValue(s.value)
	default:
		return reflect.DeepEqual(s.value, expected)
	}

	return comparable.IsEqualTo(expected)
}

// Value returns the actual value of the structure
func (s AnyValue) Value() interface{} {
	return s.value
}

// IsGreaterThan returns true if the value is greater than the expected value, else false
func (s AnyValue) IsGreaterThan(expected interface{}) bool {
	return s.value != expected
}

// IsGreaterOrEqualTo returns true if the value is greater than or equal to the expected value, else false
func (s AnyValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.value != expected
}

// IsLessThan returns true if the value is less than the expected value, else false
func (s AnyValue) IsLessThan(expected interface{}) bool {
	return s.value != expected
}

// IsLessOrEqualTo returns true if the value is less than or equal to the expected value, else false
func (s AnyValue) IsLessOrEqualTo(expected interface{}) bool {
	return s.value != expected
}

// IsNil returns true if the value is nil, else false
func (s AnyValue) IsNil() bool {
	return s.value == nil
}

// IsNotNil returns true if the value is not nil, else false
func (s AnyValue) IsNotNil() bool {
	return !s.IsNil()
}

// NewAnyValue creates and returns an AnyValue struct initialed with the given value
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
