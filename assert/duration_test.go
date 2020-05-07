package assert

import (
	"testing"
	"time"
)

func TestAssertableDuration_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Duration
		expected   time.Duration
		shouldFail bool
	}{
		{
			name:       "should assert not equal time.Durations",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(2) * time.Second,
			shouldFail: true,
		},
		{
			name:       "should assert equal time.Durations",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(1) * time.Second,
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatDuration(tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableDuration_IsNotEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Duration
		expected   time.Duration
		shouldFail bool
	}{
		{
			name:       "should assert not equal time.Durations",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(2) * time.Second,
			shouldFail: false,
		},
		{
			name:       "should assert equal time.Durations",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(1) * time.Second,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatDuration(tt.actual).IsNotEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableDuration_IsShorterThan(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Duration
		expected   time.Duration
		shouldFail bool
	}{
		{
			name:       "should succeed if shorter than",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(1) * time.Minute,
			shouldFail: false,
		},
		{
			name:       "should fail if not shorter than",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(1) * time.Second,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatDuration(tt.actual).IsShorterThan(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableDuration_IsLongerThan(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Duration
		expected   time.Duration
		shouldFail bool
	}{
		{
			name:       "should succeed if longer than",
			actual:     time.Duration(1) * time.Minute,
			expected:   time.Duration(1) * time.Second,
			shouldFail: false,
		},
		{
			name:       "should fail if not longer than",
			actual:     time.Duration(1) * time.Second,
			expected:   time.Duration(1) * time.Minute,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatDuration(tt.actual).IsLongerThan(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
