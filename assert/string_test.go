package assert

import (
	"testing"
)

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
		stringOpts []StringOpt
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
		{
			name:       "should fail if doesn't contain the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "HAPPY MAN",
			shouldFail: true,
		},
		{
			name:       "should succeed if ignore case is enabled",
			actual:     "I'm a happy man",
			substring:  "HAPPY MAN",
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringCase()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual, tt.stringOpts...).Contains(tt.substring)
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
		stringOpts []StringOpt
	}{
		{
			name:       "should succeed if it only contains the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "I'm a happy man",
			shouldFail: false,
		},
		{
			name:       "should succeed if it only contains the expected sub-string with ignoring case",
			actual:     "I'm a happy man",
			substring:  "I'M a HAPPY man",
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringCase()},
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
			ft.AssertThatString(tt.actual, tt.stringOpts...).ContainsOnly(tt.substring)
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
		stringOpts []StringOpt
	}{
		{
			name:       "should succeed if it does not contain the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "sad man",
			shouldFail: false,
		},
		{
			name:       "should fail if it contains the expected sub-string with ignoring case",
			actual:     "I'm a happy man",
			substring:  "I'M a HAPPY man",
			shouldFail: true,
			stringOpts: []StringOpt{IgnoringCase()},
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
			ft.AssertThatString(tt.actual, tt.stringOpts...).DoesNotContain(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_StartsWith(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
		stringOpts []StringOpt
	}{
		{
			name:       "should succeed it starts with substring",
			actual:     "I'm a happy man",
			substring:  "I'm a",
			shouldFail: false,
		},
		{
			name:       "should fail it starts with substring by not ignoring case",
			actual:     "I'm a happy man",
			substring:  "I'M A HA",
			shouldFail: true,
		},
		{
			name:       "should succeed it starts with substring by ignoring case",
			actual:     "I'm a happy man",
			substring:  "I'M A HA",
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringCase()},
		},
		{
			name:       "should fail if it doesn't start with the substring",
			actual:     "I'm a happy man",
			substring:  "happy man",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual, tt.stringOpts...).StartsWith(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_EndsWith(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
		stringOpts []StringOpt
	}{
		{
			name:       "should succeed it ends with a substring",
			actual:     "I'm a happy man",
			substring:  "man",
			shouldFail: false,
		},
		{
			name:       "should succeed it ends with a substring by not ignoring case",
			actual:     "I'm a happy man",
			substring:  "MAN",
			shouldFail: true,
		},
		{
			name:       "should succeed it ends with a substring by ignoring case",
			actual:     "I'm a happy man",
			substring:  "MAN",
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringCase()},
		},
		{
			name:       "should fail if it doesn't end with the substring",
			actual:     "I'm a happy man",
			substring:  "woman",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual, tt.stringOpts...).EndsWith(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_HasSameSizeAs(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should succeed it has the same size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy orc",
			shouldFail: false,
		},
		{
			name:       "should fail if it doesn't have the same size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy woman",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).HasSameSizeAs(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_ContainsOnlyDigits(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should succeed if it only contains digits",
			actual:     "1234567890",
			shouldFail: false,
		},
		{
			name:       "should succeed if it only contains one digit",
			actual:     "4",
			shouldFail: false,
		},
		{
			name:       "should succeed if it contains huge number",
			actual:     "18446744073709551616",
			shouldFail: false,
		},
		{
			name:       "should succeed if it contains huge number and a character",
			actual:     "a18446744073709551616",
			shouldFail: true,
		},
		{
			name:       "should fail if contains a letter",
			actual:     "01e",
			shouldFail: true,
		},
		{
			name:       "should fail if contains only letters",
			actual:     "test",
			shouldFail: true,
		},
		{
			name:       "should fail if contains empty string",
			actual:     " ",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).ContainsOnlyDigits()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
