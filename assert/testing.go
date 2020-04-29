package assert

import (
	"testing"
)

type fluentT struct {
	t *testing.T
}

func NewFluentT(t *testing.T) *fluentT {
	return &fluentT{
		t: t,
	}
}

func (t fluentT) AssertThatString(actual string) AssertableString {
	return ThatString(t.t, actual)
}

func (t fluentT) AssertThatBool(actual bool) AssertableBool {
	return ThatBool(t.t, actual)
}

func (t fluentT) AssertThatInt(actual int) AssertableInt {
	return ThatInt(t.t, actual)
}

func (t fluentT) AssertThatSlice(actual interface{}, opts ...SliceOpt) AssertableSlice {
	return ThatSlice(t.t, actual, opts...)
}

func (t fluentT) AssertThatStruct(actual interface{}) AssertableStruct {
	return ThatStruct(t.t, actual)
}

func (t fluentT) AssertThat(actual interface{}) Assertable {
	return That(t.t, actual)
}
