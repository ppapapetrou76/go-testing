package assert

import "testing"

func TestAssertableString_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should assert empty string",
			actual:     "",
			shouldFail: false,
		},
		{
			name:       "should assert non-empty string",
			actual:     "non-empty",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).IsEmpty()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_IsNotEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should assert empty string",
			actual:     "",
			shouldFail: true,
		},
		{
			name:       "should assert non-empty string",
			actual:     "non-empty",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).IsNotEmpty()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		expected   string
		shouldFail bool
	}{
		{
			name:       "should assert not equal strings",
			actual:     "",
			expected:   "some-string",
			shouldFail: true,
		},
		{
			name:       "should assert equal strings",
			actual:     "some-string",
			expected:   "some-string",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_IsNotEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		expected   string
		shouldFail bool
	}{
		{
			name:       "should assert not equal strings",
			actual:     "",
			expected:   "some-string",
			shouldFail: false,
		},
		{
			name:       "should assert equal strings",
			actual:     "some-string",
			expected:   "some-string",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).IsNotEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
