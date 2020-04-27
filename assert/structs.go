package assert

import (
	"reflect"
	"testing"

	"github.com/ppapapetrou76/go-testing/core"
)

// ThatStruct returns a proper assertable structure based on the slice type
func ThatStruct(t *testing.T, actual interface{}) AssertableStruct {
	return AssertableStruct{
		actual: NewStructValue(actual, t),
		t:      t,
	}
}

type AssertableStruct struct {
	actual StructValue
	t      *testing.T
}

// IsEqualTo asserts if the expected structure is equal to the assertable structure value
// It errors the tests if the compared values (actual VS expected) are not equal
func (s AssertableStruct) IsEqualTo(expected interface{}) AssertableStruct {
	if !s.actual.IsEqualTo(expected) {
		s.t.Error(core.ShouldBeEqual(s.actual, expected))
	}
	return s
}

// IsNotEqualTo asserts if the expected structure is not equal to the assertable structure value
// It errors the tests if the compared values (actual VS expected) are equal
func (s AssertableStruct) IsNotEqualTo(expected int) AssertableStruct {
	if s.actual.IsEqualTo(expected) {
		s.t.Error(core.ShouldNotBeEqual(s.actual, expected))
	}
	return s
}

type StructValue struct {
	value          interface{}
	t              *testing.T
	excludedFields []string
}

func (s AssertableStruct) ExcludingFields(fields ...string) AssertableStruct {
	s.actual.excludedFields = fields
	return s
}

func (s StructValue) IsEqualTo(expected interface{}) bool {
	if !(reflect.ValueOf(s.value).Kind() == reflect.Struct && reflect.ValueOf(expected).Kind() == reflect.Struct) {
		return false
	}
	if reflect.TypeOf(s.value) != reflect.TypeOf(expected) {
		return false
	}
	actualStruct := reflect.ValueOf(s.value)
	expectedStruct := reflect.ValueOf(expected)

	for i := 0; i < actualStruct.NumField(); i++ {
		actualValue := actualStruct.Field(i)
		expectedValue := expectedStruct.Field(i)
		if sliceContains(s.excludedFields, reflect.ValueOf(s.value).Type().Field(i).Name) {
			continue
		}
		if actualValue.CanInterface() {
			if !reflect.DeepEqual(actualValue.Interface(), expectedValue.Interface()) {
				return false
			}
		} else {
			if !areEqualValues(actualValue, expectedValue) {
				return false
			}
		}
	}
	return true
}

func (s StructValue) IsNotEqualTo(expected interface{}) bool {
	return !s.IsEqualTo(expected)
}

func (s StructValue) Value() interface{} {
	return s.value
}

func NewStructValue(value interface{}, t *testing.T) StructValue {
	return StructValue{value: value, t: t}
}

func sliceContains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func areEqualValues(actualValue, expectedValue reflect.Value) bool {
	switch actualValue.Kind() {
	case reflect.String:
		return actualValue.String() == expectedValue.String()
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return actualValue.Int() == expectedValue.Int()
	case reflect.Bool:
		return actualValue.Bool() == expectedValue.Bool()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return actualValue.Uint() == expectedValue.Uint()
	case reflect.Float32, reflect.Float64:
		return actualValue.Float() == expectedValue.Float()
	case reflect.Array, reflect.Slice:
		if actualValue.Len() != expectedValue.Len() {
			return false
		}
		if actualValue.Len() > 0 && expectedValue.Len() > 0 {
			for i := 0; i < actualValue.Len(); i++ {
				if !areEqualValues(actualValue.Index(i), expectedValue.Index(i)) {
					return false
				}
			}
		}
		return true
	case reflect.Map:
		if actualValue.Len() != expectedValue.Len() {
			return false
		}
		if actualValue.Len() > 0 && expectedValue.Len() > 0 {
			for _, k := range actualValue.MapKeys() {
				if !expectedValue.MapIndex(k).IsValid() {
					return false
				}
				if !areEqualValues(actualValue.MapIndex(k), expectedValue.MapIndex(k)) {
					return false
				}
			}
		}
		return true
	default:
		return true
	}
}
