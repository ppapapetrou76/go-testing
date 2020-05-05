package assert

import (
	"testing"
	"time"
)

func TestAssertableTime_IsSameAs(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Time
		expected   time.Time
		shouldFail bool
	}{
		{
			name:       "should assert not same time.Times",
			actual:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: true,
		},
		{
			name:       "should assert same time.Times",
			actual:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatTime(tt.actual).IsSameAs(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableTime_IsNotSameAs(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Time
		expected   time.Time
		shouldFail bool
	}{
		{
			name:       "should assert not same time.Times",
			actual:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: false,
		},
		{
			name:       "should assert equal time.Times",
			actual:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatTime(tt.actual).IsNotTheSameAs(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableTime_IsAfter(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Time
		expected   time.Time
		shouldFail bool
	}{
		{
			name:       "should succeed if after",
			actual:     time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: false,
		},
		{
			name:       "should fail if not after",
			actual:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatTime(tt.actual).IsAfter(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableInt_IsBefore(t *testing.T) {
	tests := []struct {
		name       string
		actual     time.Time
		expected   time.Time
		shouldFail bool
	}{
		{
			name:       "should succeed if before",
			actual:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: false,
		},
		{
			name:       "should fail if not before",
			actual:     time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatTime(tt.actual).IsBefore(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
