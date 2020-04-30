package assert

// AssertableSlice defines the interface for all kinds of assertable slices
type AssertableSlice interface {
	IsEqualTo(expected interface{}) AssertableSlice
	IsNotEqualTo(expected interface{}) AssertableSlice
	IsEmpty() AssertableSlice
	IsNotEmpty() AssertableSlice
	HasSize(size int) AssertableSlice
	Contains(elements interface{}) AssertableSlice
	ContainsOnly(elements interface{}) AssertableSlice
	DoesNotContain(elements interface{}) AssertableSlice
}
