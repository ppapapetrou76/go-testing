package types

// Assertable is the basic interface for all assertable values
type Assertable interface {
	Value() interface{}
}

// Comparable is the interface for basic comparable operations
type Comparable interface {
	IsEqualTo(expected interface{}) bool
	IsNotEqualTo(expected interface{}) bool
	Assertable
}

// ExtendedComparable is the interface for advanced comparable operations
type ExtendedComparable interface {
	IsGreaterOrEqualTo(expected interface{}) bool
	IsGreaterThan(expected interface{}) bool
	IsLessThan(expected interface{}) bool
	IsLessOrEqualTo(expected interface{}) bool
	Comparable
}

// Containable is the interface for operations related to containable values such as string or slice
type Containable interface {
	Contains(elements interface{}) bool
	ContainsOnly(elements interface{}) bool
	DoesNotContain(elements interface{}) bool
	IsEmpty() bool
	IsNotEmpty() bool
	HasSize(length int) bool
	Size() int
	Comparable
}

// Nullable is the interface for operations related to nullable values such as pointers or slices
type Nullable interface {
	IsNil() bool
	IsNotNil() bool
	Assertable
}
