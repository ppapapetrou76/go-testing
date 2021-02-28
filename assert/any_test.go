package assert

import (
	"reflect"
	"testing"
)

func TestAssertable_IsNil(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert nil",
			actual:     nil,
			shouldFail: false,
		},
		{
			name:       "should assert non-empty string",
			actual:     "non-nil",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			That(test, tt.actual).IsNil()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertable_IsNotNil(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert nil",
			actual:     nil,
			shouldFail: true,
		},
		{
			name:       "should assert non-empty string",
			actual:     "non-nil",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			That(test, tt.actual).IsNotNil()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertable_IsTrue(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert a false value if it's true",
			actual:     false,
			shouldFail: true,
		},
		{
			name:       "should assert a true value if it's true",
			actual:     true,
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			That(test, tt.actual).IsTrue()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertable_IsFalse(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert a false value if it's false",
			actual:     false,
			shouldFail: false,
		},
		{
			name:       "should assert a true value if it's false",
			actual:     true,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			That(test, tt.actual).IsFalse()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableAny_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{
			name:       "should assert a string slice which is not equal to expected",
			actual:     []string{"123", "321"},
			expected:   []string{"123"},
			shouldFail: true,
		},
		{
			name:       "should assert a string slice which is equal to expected",
			actual:     []string{"123", "321"},
			expected:   []string{"123", "321"},
			shouldFail: false,
		},
		{
			name:       "should assert a string slice when the expected is of different type",
			actual:     []string{"123"},
			expected:   "123",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			That(test, tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableAny_HasTypeOf(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert the same types",
			actual:     "123",
			shouldFail: false,
		},
		{
			name:       "should assert different types",
			actual:     true,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			That(test, tt.actual).HasTypeOf(reflect.TypeOf(""))
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
