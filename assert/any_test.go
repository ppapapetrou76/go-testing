package assert

import "testing"

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
