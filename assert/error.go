package assert

import (
	"fmt"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

func ShouldBeEqual(actual types.Comparable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be equal to %+v", actual.Value(), expected)
}

func ShouldNotBeEqual(actual types.Comparable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be other than %+v", actual.Value(), expected)
}

func ShouldBeGreater(actual types.ExtendedComparable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than %+v", actual.Value(), expected)
}

func ShouldBeGreaterOrEqual(actual types.ExtendedComparable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than or equal to %+v", actual.Value(), expected)
}

func ShouldBeLessThan(actual types.ExtendedComparable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than %+v", actual.Value(), expected)
}

func ShouldBeLessOrEqual(actual types.ExtendedComparable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than or equal to %+v", actual.Value(), expected)
}

func ShouldBeEmpty(actual types.Containable) string {
	return fmt.Sprintf("assertion failed: expected %+v to be empty, but it's not", actual.Value())
}

func ShouldNotBeEmpty(actual types.Containable) string {
	return fmt.Sprintf("assertion failed: expected %+v not to be empty, but it is", actual.Value())
}

func ShouldBeNil(actual types.Nullable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be nil but it wasn't", actual.Value())
}

func ShouldNotBeNil(actual types.Nullable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be non-nil but it was", actual.Value())
}

func ShouldHaveSize(actual types.Containable, expected int) string {
	return fmt.Sprintf("assertion failed: expected size of = [%d], to be but it has size of [%d] ", actual.Size(), expected)
}

func ShouldContain(actual types.Containable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain [%+v], but it doesn't", actual.Value(), elements)
}

func ShouldContainOnly(actual types.Containable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain only [%+v], but it doesn't", actual.Value(), elements)
}

func ShouldNotContain(actual types.Containable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should not contain [%+v], but it does", actual.Value(), elements)
}
