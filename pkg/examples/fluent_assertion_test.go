package examples

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestAssertedStruct(t *testing.T) {
	asserted := assertedStruct{
		BoolField:   true,
		StringField: "I am a very proud asserted string",
		IntField:    -10,
		SliceField:  []string{"elem1", "elem2"},
	}
	expected := assertedStruct{
		BoolField:   true,
		StringField: "I am a very proud asserted string",
		IntField:    20,
		SliceField:  []string{"elem1"},
	}

	ft := assert.NewFluentT(t)
	ft.AssertThatBool(asserted.BoolField).
		IsTrue().
		IsEqualTo(true).
		IsNotEqualTo(false)

	ft.AssertThatString(asserted.StringField).
		IsEqualTo("I am a very proud asserted string").
		IsNotEmpty().
		IsNotEqualTo("I'm very proud too")

	ft.AssertThatInt(asserted.IntField).
		IsEqualTo(-10).
		IsNotEqualTo(10).
		IsGreaterThan(-20).
		IsGreaterThanOrEqualTo(-20).
		IsGreaterThanOrEqualTo(-10).
		IsLessThan(0).
		IsLessThanOrEqualTo(0).
		IsLessThanOrEqualTo(-10)

	ft.AssertThatSlice(asserted.SliceField).
		HasSize(2).
		Contains("elem1").
		ContainsOnly([]string{"elem1", "elem2"}).
		DoesNotContain("elem3")

	ft.AssertThatStruct([]*assertedStruct{&asserted, &asserted}).IsNotEqualTo([]*assertedStruct{&expected})
}
