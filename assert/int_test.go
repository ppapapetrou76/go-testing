package assert

import "testing"

func TestAssertableInt_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     int
		expected   int
		shouldFail bool
	}{
		{
			name:       "should assert not equal ints",
			actual:     -10,
			expected:   10,
			shouldFail: true,
		},
		{
			name:       "should assert equal ints",
			actual:     -10,
			expected:   -10,
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatInt(tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableInt_IsNotEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     int
		expected   int
		shouldFail bool
	}{
		{
			name:       "should assert not equal ints",
			actual:     -10,
			expected:   10,
			shouldFail: false,
		},
		{
			name:       "should assert equal ints",
			actual:     -10,
			expected:   -10,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatInt(tt.actual).IsNotEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableInt_IsGreaterThan(t *testing.T) {
	tests := []struct {
		name       string
		actual     int
		expected   int
		shouldFail bool
	}{
		{
			name:       "should succeed if great than",
			actual:     -10,
			expected:   -20,
			shouldFail: false,
		},
		{
			name:       "should fail if not greater than",
			actual:     -10,
			expected:   -10,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatInt(tt.actual).IsGreaterThan(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableInt_IsGreaterThanOrEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     int
		expected   int
		shouldFail bool
	}{
		{
			name:       "should succeed if great than",
			actual:     -10,
			expected:   -20,
			shouldFail: false,
		},
		{
			name:       "should succeed if equal",
			actual:     -10,
			expected:   -10,
			shouldFail: false,
		},
		{
			name:       "should fail if not greater than",
			actual:     -10,
			expected:   0,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatInt(tt.actual).IsGreaterThanOrEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableInt_IsLessThan(t *testing.T) {
	tests := []struct {
		name       string
		actual     int
		expected   int
		shouldFail bool
	}{
		{
			name:       "should succeed if less than",
			actual:     -10,
			expected:   0,
			shouldFail: false,
		},
		{
			name:       "should fail if not less than",
			actual:     -10,
			expected:   -10,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatInt(tt.actual).IsLessThan(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableInt_IsLessThanOrEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     int
		expected   int
		shouldFail bool
	}{
		{
			name:       "should succeed if less than",
			actual:     -10,
			expected:   0,
			shouldFail: false,
		},
		{
			name:       "should succeed if equal",
			actual:     -10,
			expected:   -10,
			shouldFail: false,
		},
		{
			name:       "should fail if not less than",
			actual:     -10,
			expected:   -20,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			NewFluentT(test).AssertThatInt(tt.actual).IsLessThanOrEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
