package assert

import "testing"

type randomStruct struct {
	Field1  string
	Field2  int
	Struct1 struct {
		valueStruct1 string
	}
}

type randomStructWithPrivateFields struct {
	field1  string
	field2  int
	field3  bool
	field4  float32
	field5  uint16
	field6  []string
	map1    map[string]int
	struct1 struct {
		valueStruct1 string
	}
}

func TestAssertableStruct(t *testing.T) {
	tests := []struct {
		name            string
		actual          interface{}
		expected        interface{}
		shouldFail      bool
		excludingFields []string
	}{
		{
			name: "should assert structures of different types",
			actual: randomStruct{
				Field1:  "value1",
				Field2:  10,
				Struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected:   struct{}{},
			shouldFail: true,
		},
		{
			name: "should assert different kinds",
			actual: randomStruct{
				Field1:  "value1",
				Field2:  10,
				Struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected:   "not-a-struct",
			shouldFail: true,
		},
		{
			name: "should assert not equal structures",
			actual: randomStruct{
				Field1:  "value1",
				Field2:  10,
				Struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected: randomStruct{
				Field1: "value1",
				Field2: 10,
			},
			shouldFail: true,
		},
		{
			name: "should assert equal structures",
			actual: randomStruct{
				Field1:  "value1",
				Field2:  10,
				Struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected: randomStruct{
				Field1:  "value1",
				Field2:  10,
				Struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
		},
		{
			name: "should assert equal structures excluding a field",
			actual: randomStruct{
				Field1: "string1",
				Field2: 10,
			},
			expected: randomStruct{
				Field1: "Field1",
				Field2: 10,
			},
			excludingFields: []string{"Field1"},
		},
		{
			name: "should assert not equal structures with non-exported fields - ignoring struct fields",
			actual: randomStructWithPrivateFields{
				field1:  "value1",
				field2:  5,
				struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected: randomStructWithPrivateFields{
				field1: "value2",
				field2: 5,
			},
			shouldFail: true,
		},
		{
			name: "should assert equal structures with non-exported fields - ignoring struct fields",
			actual: randomStructWithPrivateFields{
				field1:  "value1",
				field2:  10,
				struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
			},
			shouldFail: false,
		},
		{
			name: "should assert equal structures with non-exported fields",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: -10,
				field3: true,
				field4: 2.4,
				field5: 20,
				field6: []string{"123"},
				map1: map[string]int{
					"1": 1,
					"2": 2,
				},
				struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: -10,
				field3: true,
				field4: 2.4,
				field5: 20,
				field6: []string{"123"},
				map1: map[string]int{
					"1": 1,
					"2": 2,
				},
				struct1: struct{ valueStruct1 string }{valueStruct1: "!23"},
			},
		},
		{
			name: "should assert equal structures excluding a field with non-exported fields",
			actual: randomStructWithPrivateFields{
				field1: "string1",
				field2: 10,
			},
			expected: randomStructWithPrivateFields{
				field1: "Field1",
				field2: 10,
			},
			excludingFields: []string{"field1"},
		},
		{
			name: "should assert not equal structures bool",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field3: false,
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field3: true,
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures float",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field4: 10.2,
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field4: 10.111,
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures uint",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field5: 10,
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field5: 3,
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures slice",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field6: []string{"123"},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field6: []string{"321"},
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures slice different size",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				field6: []string{"123"},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures map",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				map1: map[string]int{
					"1": 1,
					"2": 2,
				},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				map1: map[string]int{
					"1": 1,
					"3": 3,
				},
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures map",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				map1: map[string]int{
					"1": 1,
					"2": 2,
				},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				map1: map[string]int{
					"1": 1,
					"2": 3,
				},
			},
			shouldFail: true,
		},
		{
			name: "should assert not equal structures map",
			actual: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				map1: map[string]int{
					"1": 1,
					"2": 2,
				},
			},
			expected: randomStructWithPrivateFields{
				field1: "value1",
				field2: 10,
				map1: map[string]int{
					"1": 1,
				},
			},
			shouldFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &testing.T{}
			ThatStruct(test, tt.actual).ExcludingFields(tt.excludingFields...).IsEqualTo(tt.expected)
			ThatBool(t, test.Failed()).IsEqualTo(tt.shouldFail)
		})
	}
}
