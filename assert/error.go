package assert

import (
	"fmt"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

func shouldBeEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be equal to %+v", actual.Value(), expected)
}

func shouldNotBeEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be other than %+v", actual.Value(), expected)
}

func shouldBeGreater(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than %+v", actual.Value(), expected)
}

func shouldBeGreaterOrEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than or equal to %+v", actual.Value(), expected)
}

func shouldBeLessThan(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than %+v", actual.Value(), expected)
}

func shouldBeLessOrEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than or equal to %+v", actual.Value(), expected)
}

func shouldBeEmpty(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v to be empty, but it's not", actual.Value())
}

func shouldNotBeEmpty(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v not to be empty, but it is", actual.Value())
}

func shouldBeNil(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be nil but it wasn't", actual.Value())
}

func shouldNotBeNil(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be non-nil but it was", actual.Value())
}

func shouldHaveSize(actual types.Containable, expected int) string {
	return fmt.Sprintf("assertion failed: expected size of = [%d], to be but it has size of [%d] ", actual.Size(), expected)
}

func shouldContain(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain [%+v], but it doesn't", actual.Value(), elements)
}

func shouldContainOnly(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain only [%+v], but it doesn't", actual.Value(), elements)
}

func shouldNotContain(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should not contain [%+v], but it does", actual.Value(), elements)
}
