[![Codacy Badge](https://api.codacy.com/project/badge/Grade/70aaf3cfcd9d46f08ba1de5eb4156577)](https://app.codacy.com/manual/ppapapetrou76/go-testing?utm_source=github.com&utm_medium=referral&utm_content=ppapapetrou76/go-testing&utm_campaign=Badge_Grade_Dashboard)
[![Fluent Go Testing](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://app.circleci.com/pipelines/github/ppapapetrou76/go-testing?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ppapapetrou76/go-testing)](https://goreportcard.com/report/github.com/ppapapetrou76/go-testing)
[![GoDoc](https://godoc.org/github.com/ppapapetrou76/go-testing?status.svg)](https://pkg.go.dev/github.com/ppapapetrou76/go-testing)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ppapapetrou76_go-testing&metric=alert_status)](https://sonarcloud.io/dashboard?id=ppapapetrou76_go-testing)
[![codecov](https://codecov.io/gh/ppapapetrou76/go-testing/branch/master/graph/badge.svg)](https://codecov.io/gh/ppapapetrou76/go-testing)

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
## Supported
For the following types basic assertions are supported
  * int
  * uint
  * bool
  * string
  * slice
  * structure
  * map

## To be implemented soon
  * time.Duration (basic assertions)
  * time.Time (basic assertions)
  * moar string assertions
    * [x] ContainsOnlyWhitespaces
    * [x] ContainsWhitespaces
    * [x] DoesNotContainAnyWhitespaces
    * [x] DoesNotContainOnlyWhitespaces
    * [x] DoesNotEndWith
    * [x] DoesNotStartWith
    * [x] EndsWith
    * [ ] HasLineCount
    * [x] HasSameSizeAs
    * [ ] HasSizeBetween
    * [ ] HasSizeGreaterThan
    * [ ] HasSizeGreaterThanOrEqualTo
    * [x] HasSizeLessThan
    * [ ] HasSizeLessThanOrEqualTo
    * [x] IsEqualToIgnoringCase
    * [ ] IsEqualToIgnoringNewLines
    * [x] IsEqualToIgnoringWhitespace
    * [x] IsLowerCase
    * [x] IsNotEqualToIgnoringCase
    * [x] IsNotEqualToIgnoringWhitespace
    * [ ] IsSubstringOf
    * [x] IsUpperCase
