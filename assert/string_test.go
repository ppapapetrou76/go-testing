package assert

import (
	"strings"
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
		stringOpts []StringOpt
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
		{
			name:       "should assert equal strings ignoring white spaces",
			actual:     "somestring with white spaces",
			expected:   "some stringwith white spaces",
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringWhiteSpaces()},
		},
		{
			name: "should assert equal strings ignoring new lines",
			actual: `some string 
with white spaces`,
			expected: `some string with 
white spaces`,
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringNewLines()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual, tt.stringOpts...).IsEqualTo(tt.expected)
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
		StringOpt
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
		{
			name:       "should assert equal strings ignoring white spaces",
			actual:     "somestring with white spaces",
			expected:   "some stringwith white spaces",
			shouldFail: true,
			StringOpt:  IgnoringWhiteSpaces(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual, tt.StringOpt).IsNotEqualTo(tt.expected)
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

func TestAssertableString_ContainsIgnoringCase(t *testing.T) {
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
			name:       "should fail if doesn't contain only the expected sub-string",
			actual:     "I'm a happy man",
			substring:  " happy man",
			shouldFail: false,
		},
		{
			name:       "should succeed if it only contains the expected sub-string with ignoring case",
			actual:     "I'm a happy man",
			substring:  "a HAPPY man",
			shouldFail: false,
		},
		{
			name:       "should fail if doesn't contain only the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "sad man",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual, tt.stringOpts...).ContainsIgnoringCase(tt.substring)
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

func TestAssertableString_ContainsOnlyOnce(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
		stringOpts []StringOpt
	}{
		{
			name:       "should succeed if it contains the expected sub-string only once",
			actual:     "I'm a happy man",
			substring:  "happy man",
			shouldFail: false,
		},
		{
			name:       "should succeed if it contains the expected sub-string with ignoring case only once",
			actual:     "I'm a happy MAN",
			substring:  "HAPPY man",
			shouldFail: false,
			stringOpts: []StringOpt{IgnoringCase()},
		},
		{
			name:       "should fail if doesn't contain the expected sub-string",
			actual:     "I'm a happy man",
			substring:  "sad man",
			shouldFail: true,
		},
		{
			name:       "should fail if it contains the expected sub-string more than once",
			actual:     "I'm a happy man, you're a happy man",
			substring:  "happy man",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual, tt.stringOpts...).ContainsOnlyOnce(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_ContainsWhitespaces(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should succeed if it contains more than one whitespace",
			actual:     "I'm a happy person!",
			shouldFail: false,
		},
		{
			name:       "should succeed if contains one whitespace",
			actual:     "I'm happy!",
			shouldFail: false,
		},
		{
			name:       "should succeed if contains one whitespace at the end of the string",
			actual:     "Happy! ",
			shouldFail: false,
		},
		{
			name:       "should succeed if contains one whitespace at the beginning of the string",
			actual:     " Happy!",
			shouldFail: false,
		},
		{
			name:       "should fail if it contains no whitespace",
			actual:     "Happy!",
			shouldFail: true,
		},
		{
			name:       "should fail if it contains no whitespace and special characters",
			actual:     "Happy!@#%^&*(){}",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).ContainsWhitespaces()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_DoesNotContainAnyWhitespaces(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should succeed if it contains no whitespace",
			actual:     "Happy!",
			shouldFail: false,
		},
		{
			name:       "should succeed if it contains no whitespace and special characters",
			actual:     "Happy!@#%^&*(){}",
			shouldFail: false,
		},
		{
			name:       "should fail if it contains more than one whitespace",
			actual:     "I'm a happy person!",
			shouldFail: true,
		},
		{
			name:       "should fail if contains one whitespace",
			actual:     "I'm happy!",
			shouldFail: true,
		},
		{
			name:       "should fail if contains one whitespace at the end of the string",
			actual:     "Happy! ",
			shouldFail: true,
		},
		{
			name:       "should fail if contains one whitespace at the beginning of the string",
			actual:     " Happy!",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).DoesNotContainAnyWhitespaces()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_ContainsOnlyWhitespaces(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should succeed if it contains one whitespace",
			actual:     " ",
			shouldFail: false,
		},
		{
			name:       "should succeed if contains many whitespaces",
			actual:     "       ",
			shouldFail: false,
		},
		{
			name:       "should fail if has one letter at the beginning of the string",
			actual:     "a   ",
			shouldFail: true,
		},
		{
			name:       "should fail if has one letter at the middle of the string",
			actual:     "   a   ",
			shouldFail: true,
		},
		{
			name:       "should fail if has one letter at the end of the string",
			actual:     "   a",
			shouldFail: true,
		},
		{
			name:       "should fail if has no whitespace",
			actual:     "abcDEF123",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).ContainsOnlyWhitespaces()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_DoesNotContainOnlyWhitespaces(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should succeed if has one letter at the beginning of the string",
			actual:     "a   ",
			shouldFail: false,
		},
		{
			name:       "should succeed if has one letter at the middle of the string",
			actual:     "   a   ",
			shouldFail: false,
		},
		{
			name:       "should succeed if has one letter at the end of the string",
			actual:     "   a",
			shouldFail: false,
		},
		{
			name:       "should succeed if has no whitespace",
			actual:     "abcDEF123",
			shouldFail: false,
		},
		{
			name:       "should fail if it contains one whitespace",
			actual:     " ",
			shouldFail: true,
		},
		{
			name:       "should fail if contains many whitespaces",
			actual:     "       ",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).DoesNotContainOnlyWhitespaces()
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

func TestAssertableString_HasSizeBetween(t *testing.T) {
	tests := []struct {
		name        string
		actual      string
		shortString string
		longString  string
		shouldFail  bool
	}{
		{
			name:        "should fail if it has greater size than the longString",
			actual:      "I'm not a happy big man",
			shortString: "I'm a happy orc",
			longString:  "I'm not a happy orc",
			shouldFail:  true,
		},
		{
			name:        "should fail if it has the same size as the longString",
			actual:      "I'm a happy big man",
			shortString: "I'm a happy orc",
			longString:  "I'm not a happy orc",
			shouldFail:  true,
		},
		{
			name:        "should fail if it has less size than the shortString",
			actual:      "I'm a orc",
			shortString: "I'm a happy orc",
			longString:  "I'm not a happy orc",
			shouldFail:  true,
		},
		{
			name:        "should fail if it has the same size as the shortString",
			actual:      "I'm a happy man",
			shortString: "I'm a happy orc",
			longString:  "I'm not a happy orc",
			shouldFail:  true,
		},
		{
			name:        "should succeed if it has size between shortString and longString",
			actual:      "I'm a happy big man",
			shortString: "I'm a happy orc",
			longString:  "I'm not a happy big orc",
			shouldFail:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).HasSizeBetween(tt.shortString, tt.longString)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_HasSizeGreaterThan(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should fail if it has the same size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy orc",
			shouldFail: true,
		},
		{
			name:       "should fail if it has less size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy woman",
			shouldFail: true,
		},
		{
			name:       "should succeed if it has bigger size",
			actual:     "I'm a happy big man",
			substring:  "I'm a happy orc",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).HasSizeGreaterThan(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_HasSizeGreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should fail if it has less size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy woman",
			shouldFail: true,
		},
		{
			name:       "should succeed if it has bigger size",
			actual:     "I'm a happy big man",
			substring:  "I'm a happy orc",
			shouldFail: false,
		},
		{
			name:       "should succeed if it has the same size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy orc",
			shouldFail: false,
		},
		{
			name:       "should succeed if it has the same size and empty",
			actual:     "",
			substring:  "",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).HasSizeGreaterThanOrEqualTo(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_HasSizeLessThan(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should fail if it has the same size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy orc",
			shouldFail: true,
		},
		{
			name:       "should fail if it has bigger size",
			actual:     "I'm a happy big man",
			substring:  "I'm a happy orc",
			shouldFail: true,
		},
		{
			name:       "should succeed if it has less size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy woman",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).HasSizeLessThan(tt.substring)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_HasSizeLessThanOrEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substring  string
		shouldFail bool
	}{
		{
			name:       "should fail if it has bigger size",
			actual:     "I'm a happy big man",
			substring:  "I'm a happy orc",
			shouldFail: true,
		},
		{
			name:       "should succeed if it has less size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy woman",
			shouldFail: false,
		},
		{
			name:       "should succeed if it has the same size",
			actual:     "I'm a happy man",
			substring:  "I'm a happy orc",
			shouldFail: false,
		},
		{
			name:       "should succeed if it has the same size and empty",
			actual:     "",
			substring:  "",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ft := NewFluentT(test)
			ft.AssertThatString(tt.actual).HasSizeLessThanOrEqualTo(tt.substring)
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

func TestAssertableString_DoesNotStartWith(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substr     string
		shouldFail bool
	}{
		{
			name:       "should assert string that does not start with",
			actual:     "My name is Bond",
			substr:     "Je m'appelle Bond",
			shouldFail: false,
		},
		{
			name:       "should assert string that starts with",
			actual:     "My name is Bond",
			substr:     "My name",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).DoesNotStartWith(tt.substr)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_DoesNotEndWith(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		substr     string
		shouldFail bool
	}{
		{
			name:       "should assert string that ends with",
			actual:     "My name is Bond",
			substr:     "Bond",
			shouldFail: true,
		},
		{
			name:       "should assert string that does not end with",
			actual:     "My name is Bond",
			substr:     "My name",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).DoesNotEndWith(tt.substr)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_IsLowerCase(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should assert string that is not lower case",
			actual:     "My name is Bond",
			shouldFail: true,
		},
		{
			name:       "should assert string that is lower case",
			actual:     "my name is bond",
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).IsLowerCase()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableString_IsUpperCase(t *testing.T) {
	tests := []struct {
		name       string
		actual     string
		shouldFail bool
	}{
		{
			name:       "should assert string that is not upper case",
			actual:     "My name is Bond",
			shouldFail: true,
		},
		{
			name:       "should assert string that is upper case",
			actual:     strings.ToUpper("my name is bond"),
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatString(test, tt.actual).IsUpperCase()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
