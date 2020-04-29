package types


type Assertable interface {
	Value() interface{}
}

type Comparable interface {
	IsEqualTo(expected interface{}) bool
	IsNotEqualTo(expected interface{}) bool
	Assertable
}

type ExtendedComparable interface {
	IsGreaterOrEqualTo(expected interface{}) bool
	IsGreaterThan(expected interface{}) bool
	IsLessThan(expected interface{}) bool
	IsLessOrEqualTo(expected interface{}) bool
	Comparable
}

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

type Nullable interface {
	IsNil() bool
	IsNotNil() bool
	Assertable
}