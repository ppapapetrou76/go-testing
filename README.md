[![Fluent Go Testing](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://app.circleci.com/pipelines/github/ppapapetrou76/go-testing?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ppapapetrou76/go-testing)](https://goreportcard.com/report/github.com/ppapapetrou76/go-testing)
[![Coverage Status](https://coveralls.io/repos/github/ppapapetrou76/go-testing/badge.svg?branch=master)](https://coveralls.io/github/ppapapetrou76/go-testing?branch=master)

# go-testing
GoLang fluent API for test assertions

## Quick Start

```go
import (
  "testing"

  "github.com/ppapapetrou76/go-testing"
)

type assertedStruct struct {
	boolField   bool
	stringField string
	intField    int
	sliceField  []string
}


func TestAssertedStruct(t *testing.T) {
	assertedStruct := assertedStruct{}
	assertedStruct.boolField = true
	assertedStruct.stringField = "I am a very proud asserted string"
	assertedStruct.intField = -10
	assertedStruct.sliceField = []string{"elem1", "elem2"}

	ft := assert.NewFluentT(t)
	ft.AssertThatBool(assertedStruct.boolField).
		IsTrue().
		IsEqualTo(true).
		IsNotEqualTo(false)

	ft.AssertThatString(assertedStruct.stringField).
		IsEqualTo("I am a very proud asserted string").
		IsNotEmpty().
		IsNotEqualTo("I'm very proud too")

	ft.AssertThatInt(assertedStruct.intField).
		IsEqualTo(-10).
		IsNotEqualTo(10).
		IsGreaterThan(-20).
		IsGreaterThanOrEqualTo(-20).
		IsGreaterThanOrEqualTo(-10).
		IsLessThan(0).
		IsLessThanOrEqualTo(0).
		IsLessThanOrEqualTo(-10)

	ft.AssertThatSlice(assertedStruct.sliceField).
		HasSize(2).
		Contains("elem1").
		ContainsOnly([]string{"elem1", "elem2"}).
		DoesNotContain("elem3")
}

```
