package examples

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestAssertedStruct(t *testing.T) {
	assertedStruct := AssertedStruct{}
	assertedStruct.SetBoolField(true)
	assertedStruct.SetStringField("I am a very proud asserted string")
	assertedStruct.SetIntField(-10)
	assertedStruct.SetSliceField([]string{"elem1", "elem2"})

	ft := assert.NewFluentT(t)
	ft.AssertThatBool(assertedStruct.BoolField()).
		IsTrue().
		IsEqualTo(true).
		IsNotEqualTo(false)

	ft.AssertThatString(assertedStruct.StringField()).
		IsEqualTo("I am a very proud asserted string").
		IsNotEmpty().
		IsNotEqualTo("I'm very proud too")

	ft.AssertThatInt(assertedStruct.IntField()).
		IsEqualTo(-10).
		IsNotEqualTo(10).
		IsGreaterThan(-20).
		IsGreaterThanOrEqualTo(-20).
		IsGreaterThanOrEqualTo(-10).
		IsLessThan(0).
		IsLessThanOrEqualTo(0).
		IsLessThanOrEqualTo(-10)

	ft.AssertThatSlice(assertedStruct.SliceField()).
		HasSize(2).
		Contains("elem1").
		ContainsOnly([]string{"elem1", "elem2"}).
		DoesNotContain("elem3")
}
