package assert

import "testing"

func TestAssertableSlice_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     []string
		shouldFail bool
	}{
		{
			name:       "should assert empty string slice",
			actual:     []string{"element"},
			shouldFail: true,
		},
		{
			name:       "should assert non-empty string slice",
			actual:     []string{},
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).IsEmpty()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableSlice_IsNotEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     []string
		shouldFail bool
	}{
		{
			name:       "should assert empty string slice",
			actual:     []string{"element"},
			shouldFail: false,
		},
		{
			name:       "should assert non-empty string slice",
			actual:     []string{},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).IsNotEmpty()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableSlice_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     []string
		expected   []string
		shouldFail bool
	}{
		{
			name:       "should assert not equal string slices",
			actual:     []string{"element"},
			expected:   []string{"element"},
			shouldFail: false,
		},
		{
			name:       "should assert equal strings",
			actual:     []string{"element"},
			expected:   []string{"element-other"},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableSlice_IsNotEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     []string
		expected   []string
		shouldFail bool
	}{
		{
			name:       "should assert not equal string slices",
			actual:     []string{"element"},
			expected:   []string{"element"},
			shouldFail: true,
		},
		{
			name:       "should assert equal strings",
			actual:     []string{"element"},
			expected:   []string{"element-other"},
			shouldFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).IsNotEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableStringSlice_HasSize(t *testing.T) {
	tests := []struct {
		name         string
		actual       []string
		expectedSize int
		shouldFail   bool
	}{
		{
			name:         "should succeed if it has the expected size",
			actual:       []string{"element", "element2"},
			expectedSize: 2,
			shouldFail:   false,
		},
		{
			name:         "should fail if it has not the expected size",
			actual:       []string{"element", "element2"},
			expectedSize: 1,
			shouldFail:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).HasSize(tt.expectedSize)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableStringSlice_Contains(t *testing.T) {
	tests := []struct {
		name              string
		actual            []string
		elementsToContain interface{}
		shouldFail        bool
	}{
		{
			name:              "should succeed if it contains the expected elements",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: []string{"element", "element2"},
			shouldFail:        false,
		},
		{
			name:              "should succeed if it contains the expected element",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: "element",
			shouldFail:        false,
		},
		{
			name:              "should fail if does not contain the expected elements",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: []string{"element", "element4"},
			shouldFail:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).Contains(tt.elementsToContain)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableStringSlice_ContainsOnly(t *testing.T) {
	tests := []struct {
		name              string
		actual            []string
		elementsToContain interface{}
		shouldFail        bool
	}{
		{
			name:              "should succeed if it contains only the expected elements",
			actual:            []string{"element", "element2"},
			elementsToContain: []string{"element", "element2"},
			shouldFail:        false,
		},
		{
			name:              "should fail if it contains the elements but not only them",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: []string{"element", "element2"},
			shouldFail:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).ContainsOnly(tt.elementsToContain)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableStringSlice_DoesNotContain(t *testing.T) {
	tests := []struct {
		name              string
		actual            []string
		elementsToContain interface{}
		shouldFail        bool
	}{
		{
			name:              "should succeed if it doesn't contain the expected elements",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: []string{"element", "element10"},
			shouldFail:        false,
		},
		{
			name:              "should succeed if it doesn't the expected element",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: "element10",
			shouldFail:        false,
		},
		{
			name:              "should fail if contains the expected elements",
			actual:            []string{"element", "element2", "element3"},
			elementsToContain: []string{"element", "element2"},
			shouldFail:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).DoesNotContain(tt.elementsToContain)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
