package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/types"
)

func TestAssertableMap_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert empty string map",
			actual:     map[string]int{},
			shouldFail: false,
		},
		{
			name:       "should assert non-empty string map",
			actual:     map[string]int{"1": 1},
			shouldFail: true,
		},
		{
			name:       "should assert a non-map type",
			actual:     200,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).IsEmpty()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_IsNotEmpty(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		shouldFail bool
	}{
		{
			name:       "should assert empty string map",
			actual:     map[string]int{},
			shouldFail: true,
		},
		{
			name:   "should assert non-empty string map",
			actual: map[string]int{"1": 1},
		},
		{
			name:       "should assert a non-map type",
			actual:     200,
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).IsNotEmpty()
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_IsEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{
			name:       "should assert equal string maps",
			actual:     map[string]int{"1": 1},
			expected:   map[string]int{"1": 1},
			shouldFail: false,
		},
		{
			name:       "should assert wrong type",
			actual:     1100,
			expected:   map[string]int{"1": 1, "2": 2},
			shouldFail: true,
		},
		{
			name:       "should assert not equal length string maps",
			actual:     map[string]int{"1": 1},
			expected:   map[string]int{"1": 1, "2": 2},
			shouldFail: true,
		},
		{
			name:       "should assert different types",
			actual:     map[string]int{"1": 1},
			expected:   "123",
			shouldFail: true,
		},
		{
			name:       "should assert equal length but different keys string maps",
			actual:     map[string]int{"1": 1},
			expected:   map[string]int{"2": 1},
			shouldFail: true,
		},
		{
			name:       "should assert equal length but different values string maps",
			actual:     map[string]int{"1": 1},
			expected:   map[string]int{"1": 2},
			shouldFail: true,
		},
		{
			name:       "should assert maps of different key types",
			actual:     map[string]int{"1": 1},
			expected:   map[int]string{1: "1"},
			shouldFail: true,
		},
		{
			name:       "should assert maps of different value types",
			actual:     map[string]int{"1": 1},
			expected:   map[string]string{"1": "1"},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_IsNotEqualTo(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{
			name:       "should assert equal string maps",
			actual:     map[string]interface{}{"1": 1},
			expected:   map[string]interface{}{"1": 1},
			shouldFail: true,
		},
		{
			name:       "should assert wrong type",
			actual:     1100,
			expected:   map[string]int{"1": 1, "2": 2},
			shouldFail: true,
		},
		{
			name:     "should assert not equal length string maps",
			actual:   map[string]interface{}{"1": 1},
			expected: map[string]interface{}{"1": 1, "2": 2},
		},
		{
			name:     "should assert different types",
			actual:   map[string]interface{}{"1": 1},
			expected: "123",
		},
		{
			name:     "should assert equal length but different keys string maps",
			actual:   map[string]interface{}{"1": 1},
			expected: map[string]interface{}{"2": 1},
		},
		{
			name:     "should assert equal length but different values string maps",
			actual:   map[string]interface{}{"1": 1},
			expected: map[string]int{"1": 2},
		},
		{
			name:     "should assert maps of different key types",
			actual:   map[string]interface{}{"1": 1},
			expected: map[int]string{1: "1"},
		},
		{
			name:     "should assert maps of different value types",
			actual:   map[string]interface{}{"1": 1},
			expected: map[string]string{"1": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).IsNotEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableStringMap_HasSize(t *testing.T) {
	tests := []struct {
		name         string
		actual       interface{}
		expectedSize int
		shouldFail   bool
	}{
		{
			name:         "should succeed if it has the expected size",
			actual:       map[bool]string{true: "element", false: "element2"},
			expectedSize: 2,
			shouldFail:   false,
		},
		{
			name:         "should fail if it has not the expected size",
			actual:       map[bool]string{true: "element", false: "element2"},
			expectedSize: 1,
			shouldFail:   true,
		},
		{
			name:         "should fail if it has not the expected type",
			actual:       20,
			expectedSize: 1,
			shouldFail:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasSize(tt.expectedSize)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_HasKey(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		key        interface{}
		shouldFail bool
	}{
		{
			name:       "should succeed if it has the expected key",
			actual:     map[string]interface{}{"1": 1},
			key:        "1",
			shouldFail: false,
		},
		{
			name:       "should succeed if it doesnt' have the key",
			actual:     map[string]interface{}{},
			key:        "1",
			shouldFail: true,
		},
		{
			name:       "should fail for invalid types",
			actual:     1,
			key:        "1",
			shouldFail: true,
		},
		{
			name:       "should fail for not comparable types",
			actual:     map[string]interface{}{},
			key:        []string{},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasKey(tt.key)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_HasValue(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		value      interface{}
		shouldFail bool
	}{
		{
			name:       "should succeed if it has the expected value",
			actual:     map[string]int{"1": 1},
			value:      1,
			shouldFail: false,
		},
		{
			name:       "should succeed if it doesnt' have the key",
			actual:     map[string]interface{}{},
			value:      "1",
			shouldFail: true,
		},
		{
			name:       "should fail for invalid types",
			actual:     1,
			value:      "1",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasValue(tt.value)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_HasEntry(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		entry      types.MapEntry
		shouldFail bool
	}{
		{
			name:       "should succeed if it has the expected entry",
			actual:     map[string]int{"1": 1},
			entry:      types.NewMapEntry("1", 1),
			shouldFail: false,
		},
		{
			name:       "should succeed if it doesnt' have the entry",
			actual:     map[string]int{},
			entry:      types.NewMapEntry("1", 1),
			shouldFail: true,
		},
		{
			name:       "should fail for invalid types",
			actual:     1,
			entry:      types.NewMapEntry("1", 1),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasEntry(tt.entry)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_HasNotKey(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		key        interface{}
		shouldFail bool
	}{
		{
			name:       "should succeed if it doesnt' have the key",
			actual:     map[string]interface{}{"1": 1},
			key:        "2",
			shouldFail: false,
		},
		{
			name:       "should fail if it has the key",
			actual:     map[string]interface{}{"1": 1},
			key:        "1",
			shouldFail: true,
		},
		{
			name:       "should fail for invalid types",
			actual:     1,
			key:        "1",
			shouldFail: true,
		},
		{
			name:   "should succeed for not comparable types",
			actual: map[string]interface{}{},
			key:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasNotKey(tt.key)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_HasNotValue(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		value      interface{}
		shouldFail bool
	}{
		{
			name:       "should succeed if it has not the expected value",
			actual:     map[string]interface{}{"1": 1},
			value:      2,
			shouldFail: false,
		},
		{
			name:       "should fail if it has the key",
			actual:     map[string]interface{}{"1": 1},
			value:      1,
			shouldFail: true,
		},
		{
			name:       "should fail for invalid types",
			actual:     1,
			value:      "1",
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasNotValue(tt.value)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}

func TestAssertableMap_HasNotEntry(t *testing.T) {
	tests := []struct {
		name       string
		actual     interface{}
		entry      types.MapEntry
		shouldFail bool
	}{
		{
			name:       "should fail if it has the expected entry",
			actual:     map[string]interface{}{"1": 1},
			entry:      types.NewMapEntry("1", 1),
			shouldFail: true,
		},
		{
			name:       "should succeed if it doesn't have the entry",
			actual:     map[string]interface{}{"1": 2},
			entry:      types.NewMapEntry("1", 1),
			shouldFail: false,
		},
		{
			name:       "should succeed if it doesn't have the entry",
			actual:     map[string]interface{}{"2": 1},
			entry:      types.NewMapEntry("1", 1),
			shouldFail: false,
		},
		{
			name:       "should fail for invalid types",
			actual:     1,
			entry:      types.NewMapEntry("1", 1),
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatMap(test, tt.actual).HasNotEntry(tt.entry)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
