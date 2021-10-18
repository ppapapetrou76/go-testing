package values

import (
	"reflect"
)

// SliceValue is a struct that holds a string slice value.
type SliceValue struct {
	value interface{}
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (s SliceValue) IsEqualTo(expected interface{}) bool {
	if !IsSlice(expected) || !IsSlice(s.Value()) {
		return false
	}

	actualValue := reflect.ValueOf(s.value)
	expectedValue := reflect.ValueOf(expected)
	return areSlicesEqual(actualValue, expectedValue)
}

// IsEmpty returns true if the slice is empty else false.
func (s SliceValue) IsEmpty() bool {
	return s.HasSize(0)
}

// IsNotEmpty returns true if the slice is not empty else false.
func (s SliceValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// HasSize returns true if the slice has the expected size else false.
func (s SliceValue) HasSize(length int) bool {
	return s.Size() == length
}

// Size returns the slice size.
func (s SliceValue) Size() int {
	if !IsSlice(s.Value()) {
		return 0
	}
	return reflect.ValueOf(s.Value()).Len()
}

func (s SliceValue) contains(element reflect.Value) bool {
	actualValue := reflect.ValueOf(s.Value())

	for i := 0; i < actualValue.Len(); i++ {
		if areEqualValues(actualValue.Index(i), element) {
			return true
		}
	}
	return false
}

// Contains returns true if the slice contains the expected element(s) else false.
func (s SliceValue) Contains(elements interface{}) bool {
	if !IsSlice(s.Value()) {
		return false
	}

	if !IsSlice(elements) {
		return s.contains(reflect.ValueOf(elements))
	}

	expectedValue := reflect.ValueOf(elements)
	all := true

	for i := 0; i < expectedValue.Len(); i++ {
		all = all && s.contains(expectedValue.Index(i))
	}
	return all
}

// DoesNotContain returns true if the slice does not contain the expected element(s) else false.
func (s SliceValue) DoesNotContain(elements interface{}) bool {
	return !s.Contains(elements)
}

// ContainsOnly returns true if the slice contains only the expected element(s) else false.
func (s SliceValue) ContainsOnly(elements interface{}) bool {
	return s.Contains(elements) && s.HasSize(reflect.ValueOf(elements).Len())
}

// HasUniqueElements returns true if the slice contains only unique elements else false.
func (s SliceValue) HasUniqueElements() bool {
	if !IsSlice(s.Value()) {
		return false
	}
	sliceValue := reflect.ValueOf(s.value)
	elements := map[interface{}]bool{}

	for i := 0; i < sliceValue.Len(); i++ {
		if _, ok := elements[sliceValue.Index(i).Interface()]; ok {
			return false
		}
		elements[sliceValue.Index(i).Interface()] = true
	}
	return true
}

// Value returns the actual value of the structure.
func (s SliceValue) Value() interface{} {
	return s.value
}

// NewSliceValue creates and returns a SliceValue struct initialed with the given value.
func NewSliceValue(value interface{}) SliceValue {
	return SliceValue{value: value}
}

// IsSlice returns true if the given value is a slice, else false.
func IsSlice(value interface{}) bool {
	return reflect.ValueOf(value).Kind() == reflect.Slice || reflect.ValueOf(value).Kind() == reflect.Array
}
