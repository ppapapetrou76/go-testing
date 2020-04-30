package types

import (
	"fmt"
	"reflect"
	"strings"
)

// StringSliceValue is a struct that holds a string slice value
type StringSliceValue struct {
	value []string
}

// IsEqualTo returns true if the value is equal to the expected value, else false
func (s StringSliceValue) IsEqualTo(expected interface{}) bool {
	return reflect.DeepEqual(s.value, expected)
}

// IsNotEqualTo returns true if the value is not equal to the expected value, else false
func (s StringSliceValue) IsNotEqualTo(expected interface{}) bool {
	return !s.IsEqualTo(expected)
}

// IsEmpty returns true if the slice is empty else false
func (s StringSliceValue) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty returns true if the slice is not empty else false
func (s StringSliceValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// HasSize returns true if the slice has the expected size else false
func (s StringSliceValue) HasSize(length int) bool {
	return len(s.value) == length
}

// Size returns the slice size
func (s StringSliceValue) Size() int {
	return len(s.value)
}

func (s StringSliceValue) contains(element string) bool {
	for _, item := range s.value {
		if item == element {
			return true
		}
	}
	return false
}

// Contains returns true if the slice contains the expected element(s) else false
func (s StringSliceValue) Contains(elements interface{}) bool {
	switch v := elements.(type) {
	case []string:
		var all = true
		for _, c := range v {
			all = all && s.contains(c)
		}
		return all
	case string:
		return s.contains(v)
	default:
		return false
	}
}

// DoesNotContain returns true if the slice does not contain the expected element(s) else false
func (s StringSliceValue) DoesNotContain(elements interface{}) bool {
	return !s.Contains(elements)
}

// ContainsOnly returns true if the slice contains only the expected element(s) else false
func (s StringSliceValue) ContainsOnly(elements interface{}) bool {
	switch v := elements.(type) {
	case []string:
		return s.HasSize(len(v)) && s.Contains(elements)
	case string:
		return s.HasSize(1) && s.contains(v)
	default:
		return false
	}
}

// Value returns the actual value of the structure
func (s StringSliceValue) Value() interface{} {
	return s.value
}

// NewStringSliceValue creates and returns a StringSliceValue struct initialed with the given value
func NewStringSliceValue(value interface{}) StringSliceValue {
	switch v := value.(type) {
	case []string:
		return StringSliceValue{value: v}
	default:
		panic(fmt.Sprintf("expected string slice value type but got %T type", v))
	}
}

func (s StringSliceValue) String() string {
	return strings.Join(s.value, ",")
}
