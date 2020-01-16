package core

import (
	"fmt"
)

func ShouldBeEqual(actual BasicAssertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be equal to %+v", actual.Value(), expected)
}

func ShouldNotBeEqual(actual BasicAssertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be other than %+v", actual.Value(), expected)
}

func ShouldBeGreater(actual BasicAssertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than %+v", actual.Value(), expected)
}

func ShouldBeGreaterOrEqual(actual BasicAssertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than or equal to %+v", actual.Value(), expected)
}

func ShouldBeLessThan(actual BasicAssertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than %+v", actual.Value(), expected)
}

func ShouldBeLessOrEqual(actual BasicAssertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than or equal to %+v", actual.Value(), expected)
}

func ShouldBeEmpty(actual ContainableAssertable) string {
	return fmt.Sprintf("assertion failed: expected %+v to be empty, but it's not", actual.Value())
}

func ShouldNotBeEmpty(actual ContainableAssertable) string {
	return fmt.Sprintf("assertion failed: expected %+v not to be empty, but it is", actual.Value())
}

