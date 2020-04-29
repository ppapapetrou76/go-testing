package types

import (
	"fmt"
	"reflect"
	"strings"
)

type StringSliceValue struct {
	value []string
}

func (s StringSliceValue) IsEqualTo(expected interface{}) bool {
	return reflect.DeepEqual(s.value, expected)
}

func (s StringSliceValue) IsNotEqualTo(expected interface{}) bool {
	return !s.IsEqualTo(expected)
}

func (s StringSliceValue) IsEmpty() bool {
	return len(s.value) == 0
}

func (s StringSliceValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s StringSliceValue) HasSize(length int) bool {
	return len(s.value) == length
}

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

func (s StringSliceValue) DoesNotContain(elements interface{}) bool {
	return !s.Contains(elements)
}

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

func (s StringSliceValue) Value() interface{} {
	return s.value
}

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
