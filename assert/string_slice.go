package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ppapapetrou76/go-testing/core"
)

// AssertableString is the implementation of CommonAssertable for string types
type AssertableStringSlice struct {
	T      *testing.T
	Actual []string
}

// HasSize asserts if the assertable string slice has the expected length size
// It errors the test if it doesn't have the expected size
func (a AssertableStringSlice) HasSize(size int, customMessage ...string) core.AssertableSlice {
	message := fmt.Sprintf("%s assertion failed : expected string slice length = %d, bit actual length is %d", a.T.Name(), size, len(a.Actual))
	if len(customMessage) > 0 {
		message = customMessage[0]
	}

	if len(a.Actual) != size {
		a.T.Error(message)
	}

	return a
}

// IsEmpty asserts if the assertable string slice is empty
// It errors the test if it is not empty
func (a AssertableStringSlice) IsEmpty(customMessage ...string) core.AssertableSlice {
	message := fmt.Sprintf("%s assertion failed : expected string slice to be empty but it contains %d elements", a.T.Name(), len(a.Actual))
	if len(customMessage) > 0 {
		message = customMessage[0]
	}

	if len(a.Actual) != 0 {
		a.T.Error(message)
	}

	return a
}

// IsNotEmpty asserts if the assertable string slice is not empty
// It errors the test if it is empty
func (a AssertableStringSlice) IsNotEmpty(customMessage ...string) core.AssertableSlice {
	message := fmt.Sprintf("%s assertion failed : expected string slice to has elements but it is empty", a.T.Name())
	if len(customMessage) > 0 {
		message = customMessage[0]
	}

	if len(a.Actual) == 0 {
		a.T.Error(message)
	}

	return a
}

// Contains asserts if the assertable string slice contains the given element
// It errors the test if
// - it does not contain it
// - if the slice is not of []string type
func (a AssertableStringSlice) Contains(element interface{}, customMessage ...string) core.AssertableSlice {
	message := fmt.Sprintf("%s assertion failed : expected string slice %+v to contain element %v but it's not", a.T.Name(), a, element)
	if len(customMessage) > 0 {
		message = customMessage[0]
	}

	contains := false
	switch v := element.(type) {
	case string:
		for _, e := range a.Actual {
			if e == v {
				contains = true
				break
			}
		}
	default:
		a.T.Error(message)
	}

	if !contains {
		a.T.Error(message)
	}

	return a
}

// Contains asserts if the assertable string slice does not contain the given element
// It errors the test if it contains it
func (a AssertableStringSlice) DoesNotContain(element interface{}, customMessage ...string) core.AssertableSlice {
	message := fmt.Sprintf("%s assertion failed : expected string slice %+v to not contain element %v but it is", a.T.Name(), a, element)
	if len(customMessage) > 0 {
		message = customMessage[0]
	}

	switch v := element.(type) {
	case string:
		for _, e := range a.Actual {
			if e == v {
				a.T.Error(message)
			}
		}
	default:
		return a
	}

	return a
}

// Contains asserts if the assertable string slice is equal to the expected slice
// It errors the test
// - if the slices are not equal
// - if the slice is not of []string type
func (a AssertableStringSlice) IsEqualTo(expected interface{}, customMessage ...string) core.AssertableSlice {
	message := fmt.Sprintf("%s assertion failed : expected string slice %+v to be equal to %+v but it's not", a.T.Name(), a, expected)

	switch v := expected.(type) {
	case []string:
		a.HasSize(len(v))
		if !reflect.DeepEqual(a.Actual, expected) {
			a.T.Error(message)
		}
	default:
		a.T.Error(message)
	}

	return a
}
