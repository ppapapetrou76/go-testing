package core

import (
	"testing"
)

type Assertable interface {
	IsEqualTo(expected interface{}) bool
	IsNotEqualTo(expected interface{}) bool
	Value() interface{}
}

type BasicAssertable interface {
	Assertable
	IsGreaterOrEqualTo(expected interface{}) bool
	IsGreaterThan(expected interface{}) bool
	IsLessThan(expected interface{}) bool
	IsLessOrEqualTo(expected interface{}) bool
}

type ContainableAssertable interface {
	Contains(elements ...interface{}) bool
	ContainsOnly(elements ...interface{}) bool
	IsEmpty() bool
	IsNotEmpty() bool
	Len() int
	Value() interface{}
}

type CommonContainableAssertable struct {
	T *testing.T
}

func (a CommonContainableAssertable) IsEmpty(assertable ContainableAssertable) {
	if assertable.Len() != 0 {
		a.T.Error(ShouldBeEmpty(assertable))
	}
}

func (a CommonContainableAssertable) IsNotEmpty(assertable ContainableAssertable) {
	if assertable.Len() == 0 {
		a.T.Error(ShouldNotBeEmpty(assertable))
	}
}

type CommonAssertable struct {
	T *testing.T
}

func (a CommonAssertable) IsEqualTo(assertable Assertable, expected interface{}) {
	if !assertable.IsEqualTo(expected) {
		a.T.Error(ShouldBeEqual(assertable, expected))
	}
}

func (a CommonAssertable) IsNotEqualTo(assertable Assertable, expected interface{}) {
	if assertable.IsEqualTo(expected) {
		a.T.Error(ShouldNotBeEqual(assertable, expected))
	}
}

func (a CommonAssertable) IsGreaterThan(assertable BasicAssertable, expected interface{}) {
	if assertable.IsGreaterThan(expected) {
		a.T.Error(ShouldBeGreater(assertable, expected))
	}
}

func (a CommonAssertable) IsGreaterOrEqualTo(assertable BasicAssertable, expected interface{}) {
	if assertable.IsGreaterOrEqualTo(expected) {
		a.T.Error(ShouldBeGreaterOrEqual(assertable, expected))
	}
}

func (a CommonAssertable) IsLessThan(assertable BasicAssertable, expected interface{}) {
	if assertable.IsLessThan(expected) {
		a.T.Error(ShouldBeLessThan(assertable, expected))
	}
}

func (a CommonAssertable) IsLessOrEqualTo(assertable BasicAssertable, expected interface{}) {
	if assertable.IsLessOrEqualTo(expected) {
		a.T.Error(ShouldBeLessOrEqual(assertable, expected))
	}
}

func (a CommonAssertable) IsNil(assertable BasicAssertable) {
	if assertable.Value() != nil {
		a.T.Error(ShouldBeNil(assertable))
	}
}

func (a CommonAssertable) IsNotNil(assertable BasicAssertable) {
	if assertable.Value() == nil {
		a.T.Error(ShouldNotBeNil(assertable))
	}
}
