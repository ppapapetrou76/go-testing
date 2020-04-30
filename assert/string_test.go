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

func TestAssertableString_Contains(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should succeed if it contains the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "happy man",
			shouldFail: false,
		},
		{
			name:       "should fail if doesn't contain the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "sad man",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).Contains(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_ContainsOnly(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should succeed if it only contains the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "I'm a happy man",
			shouldFail: false,
		},
		{
			name:       "should fail if doesn't contain only the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "happy man",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).ContainsOnly(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_DoesNotContain(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should succeed if it does not contain the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "sad man",
			shouldFail: false,
		},
		{
			name:       "should fail if it contains the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "happy man",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).DoesNotContain(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
