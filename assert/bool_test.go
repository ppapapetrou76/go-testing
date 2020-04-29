package assert

import "testing"

func TestAssertableBool_IsFalse(t *testing.T) {
	tests := []struct {
		name       string
		actual     bool
		shouldFail bool
	}{
		{
			name:       "should not fail if bool is false",
			actual:     false,
			shouldFail: false,
		},
		{
			name:       "should fail if bool is true",
			actual:     true,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatBool(test, tt.actual).IsFalse()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableBool_IsTrue(t *testing.T) {
	tests := []struct {
		name       string
		actual     bool
		shouldFail bool
	}{
		{
			name:       "should not fail if bool is true",
			actual:     true,
			shouldFail: false,
		},
		{
			name:       "should fail if bool is false",
			actual:     false,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatBool(test, tt.actual).IsTrue()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableBool_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     bool
		expected   interface{}
		shouldFail bool
	}{
		{
			name:       "should fail",
			actual:     true,
			expected:   false,
			shouldFail: true,
		},
		{
			name:       "should succeed",
			actual:     false,
			expected:   false,
			shouldFail: false,
		},
		{
			name:       "should fail if compared to a non-bool value",
			actual:     true,
			expected:   struct{}{},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}

			ft := NewFluentT(test)
			ft.AssertThatBool(tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableBool_IsNotEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     bool
		expected   interface{}
		shouldFail bool
	}{
		{
			name:       "should fail",
			actual:     true,
			expected:   true,
			shouldFail: true,
		},
		{
			name:       "should succeed",
			actual:     false,
			expected:   true,
			shouldFail: false,
		},
		{
			name:       "should succed if compared to a non-bool value",
			actual:     true,
			expected:   struct{}{},
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}

			ft := NewFluentT(test)
			ft.AssertThatBool(tt.actual).IsNotEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
