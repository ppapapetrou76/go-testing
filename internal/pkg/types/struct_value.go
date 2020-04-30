package types

import "reflect"

// StructValue is a struct that holds a struct value
type StructValue struct {
	value          interface{}
	ExcludedFields []string
}

// IsEqualTo returns true if the value is equal to the expected value, else false
func (s StructValue) IsEqualTo(expected interface{}) bool {
	actualElement := reflect.ValueOf(s.value)
	expectedElement := reflect.ValueOf(expected)

	if actualElement.Kind() == reflect.Ptr && expectedElement.Kind() == reflect.Ptr {
		actualElement = actualElement.Elem()
		expectedElement = expectedElement.Elem()
	}
	if !(actualElement.Kind() == reflect.Struct && expectedElement.Kind() == reflect.Struct) {
		return false
	}
	if reflect.TypeOf(s.value) != reflect.TypeOf(expected) {
		return false
	}
	for i := 0; i < actualElement.NumField(); i++ {
		actualValue := actualElement.Field(i)
		expectedValue := expectedElement.Field(i)

		if sliceContains(s.ExcludedFields, actualElement.Type().Field(i).Name) {
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

// IsNotEqualTo returns true if the value is not equal to the expected value, else false
func (s StructValue) IsNotEqualTo(expected interface{}) bool {
	return !s.IsEqualTo(expected)
}

// Value returns the actual value of the structure
func (s StructValue) Value() interface{} {
	return s.value
}

// NewStructValue creates and returns a StructValue struct initialed with the given value
func NewStructValue(value interface{}) StructValue {
	return StructValue{value: value}
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
		return areSlicesEqual(actualValue, expectedValue)
	case reflect.Map:
		return areMapsEqual(actualValue, expectedValue)
	default:
		return true
	}
}

func areMapsEqual(actualValue, expectedValue reflect.Value) bool {
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
}

func areSlicesEqual(actualValue, expectedValue reflect.Value) bool {
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
}