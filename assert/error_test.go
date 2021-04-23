package assert

import (
	"errors"
	"testing"
)

func TestAssertableError_IsNil(t *testing.T) {
	tests := []struct {
		name       string
		actual     error
		shouldFail bool
	}{
		{
			name:       "should assert nil error",
			actual:     nil,
			shouldFail: false,
		},
		{
			name:       "should assert not nil error",
			actual:     errors.New("some error"),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatError(test, tt.actual).IsNil()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableError_IsNotNil(t *testing.T) {
	tests := []struct {
		name       string
		actual     error
		shouldFail bool
	}{
		{
			name:       "should assert nil error",
			actual:     nil,
			shouldFail: true,
		},
		{
			name:       "should assert not nil error",
			actual:     errors.New("some error"),
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatError(test, tt.actual).IsNotNil()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableError_HasExactMessage(t *testing.T) {
	tests := []struct {
		name       string
		actual     error
		message    string
		shouldFail bool
	}{
		{
			name:       "should assert nil error",
			actual:     nil,
			shouldFail: true,
		},
		{
			name:       "should assert not nil error with exact message",
			actual:     errors.New("some error"),
			message:    "some error",
			shouldFail: false,
		},
		{
			name:       "should assert not nil error with different message",
			actual:     errors.New("completely different message"),
			message:    "some error",
			shouldFail: true,
		},
		{
			name:       "should assert not nil error that contains the message but not only that",
			actual:     errors.New("some error with additional information"),
			message:    "some error",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatError(test, tt.actual).HasExactMessage(tt.message)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableError_IsSameAs(t *testing.T) {
	tests := []struct {
		name       string
		actual     error
		expected   error
		shouldFail bool
	}{
		{
			name:       "should assert both nil errors",
			shouldFail: false,
		},
		{
			name:       "should assert when actual error is nil and expected is not",
			expected:   errors.New("some error"),
			shouldFail: true,
		},
		{
			name:       "should assert when actual error is not nil and expected is nil",
			actual:     errors.New("some error"),
			shouldFail: true,
		},
		{
			name:       "should assert when both errors are not nil and have the same message",
			actual:     errors.New("some error"),
			expected:   errors.New("some error"),
			shouldFail: false,
		},
		{
			name:       "should assert when both errors are not nil and have different message",
			actual:     errors.New("some error"),
			expected:   errors.New("some other error"),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatError(test, tt.actual).IsSameAs(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
