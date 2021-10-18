package assert

import "testing"

func TestAssertableSlice_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert empty string slice",
			actual:     []string{"element"},
			shouldFail: true,
		},
		{
			name:       "should assert wrong type",
			actual:     12,
			shouldFail: false,
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
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert empty string slice",
			actual:     []string{"element"},
			shouldFail: false,
		},
		{
			name:       "should assert wrong type",
			actual:     12,
			shouldFail: true,
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
	type struct1 struct {
		field1 string
		field2 bool
	}
	val1 := struct1{field1: "value1", field2: true}
	val2 := struct1{field1: "value2", field2: true}

	tests := []struct {
		name       string
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{
			name:       "should assert not equal slices of struct",
			actual:     []struct1{val1},
			expected:   []struct1{val2},
			shouldFail: true,
		},
		{
			name:       "should assert equal slices of struct",
			actual:     []struct1{val1},
			expected:   []struct1{val1},
			shouldFail: false,
		},
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
		{
			name:       "should assert slices of different sizes",
			actual:     []string{"element", "element2"},
			expected:   []string{"element-other"},
			shouldFail: true,
		},
		{
			name:       "should assert slices of different types",
			actual:     []string{"element2"},
			expected:   []int{2},
			shouldFail: true,
		},
		{
			name:       "should assert wrong type",
			actual:     2,
			expected:   []int{2},
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

func TestAssertableSlice_HasSize(t *testing.T) {
	tests := []struct {
		name         string
		actual       interface{}
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
		{
			name:         "should fail if it runs for wrong type",
			actual:       12,
			expectedSize: 1,
			shouldFail:   true,
		},
		{
			name:         "should fail if it runs for wrong type with custom message",
			actual:       12,
			expectedSize: 1,
			shouldFail:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			assertedSlice := ThatSlice(test, tt.actual, WithCustomMessage("Custom message")).HasSize(tt.expectedSize)
			ft := NewFluentT(t)
			ft.AssertThatString(assertedSlice.customMessage).IsEqualTo("Custom message")
			ft.AssertThatBool(test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableSlice_Contains(t *testing.T) {
	type testStruct struct {
		Value1 string
		Value2 int
	}

	tests := []struct {
		name              string
		actual            interface{}
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
		{
			name:              "should fail if contains runs on the wrong type",
			actual:            2,
			elementsToContain: []string{"element", "element4"},
			shouldFail:        true,
		},
		{
			name: "should succeed if contains runs on a single element not wrapped as slice",
			actual: []testStruct{
				{
					Value1: "123",
					Value2: 123,
				},
				{
					Value1: "133",
					Value2: 133,
				},
			},
			elementsToContain: testStruct{
				Value1: "123",
				Value2: 123,
			},
			shouldFail: false,
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

func TestAssertableSlice_ContainsOnly(t *testing.T) {
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

func TestAssertableSlice_DoesNotContain(t *testing.T) {
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

func TestAssertableSlice_HasUniqueElements(t *testing.T) {
	tests := []struct {
		name       string
		actual     []string
		shouldFail bool
	}{
		{
			name:       "should not fail if slice has unique elements",
			actual:     []string{"element", "element2", "element3"},
			shouldFail: false,
		},
		{
			name:       "should fail if slice has unique elements",
			actual:     []string{"element", "element", "element3"},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatSlice(test, tt.actual).HasUniqueElements()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
	test := &testing.T{}
	ThatSlice(test, []int{1, 2, 3}).HasUniqueElements()
	ThatBool(t, test.Failed()).IsEqualTo(false)
	ThatSlice(test, []int{1, 1, 3}).HasUniqueElements()
	ThatBool(t, test.Failed()).IsEqualTo(true)
}
