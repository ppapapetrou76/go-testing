package core

type AssertableSlice interface {
	Contains(element interface{}, customMessage ...string) AssertableSlice
	DoesNotContain(element interface{}, customMessage ...string) AssertableSlice
	HasSize(size int, customMessage ...string) AssertableSlice
	IsEmpty(customMessage ...string) AssertableSlice
	IsEqualTo(expected interface{}, customMessage ...string) AssertableSlice
	IsNotEmpty(customMessage ...string) AssertableSlice
}


