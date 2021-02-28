package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
	"github.com/ppapapetrou76/go-testing/types"
)

func Test_shouldBeEqual(t *testing.T) {
	type assertedStruct struct {
		BoolField   bool
		StringField string
		IntField    int
		SliceField  []string
	}
	type assertedStructUnexportedFields struct {
		boolField bool
	}

	tests := []struct {
		name            string
		actual          types.Assertable
		expected        interface{}
		expectedMessage string
	}{
		{
			name:   "should return expected message when structs have unexported fields",
			actual: values.NewStructValue(assertedStructUnexportedFields{}),
			expected: assertedStructUnexportedFields{
				boolField: true,
			},
			expectedMessage: "assertion failed:\n" +
				"expected value\t:{boolField:true}\n" +
				"actual value\t:{boolField:false}\n",
		},
		{
			name:     "should return expected message when slices are not equal",
			actual:   values.NewSliceValue([]string{"elem1", "elem2", "elem3"}),
			expected: []string{"elem1", "elem4"},
			expectedMessage: "assertion failed:\n" +
				"expected value\t:[elem1 elem4]\n" +
				"actual value\t:[elem1 elem2 elem3]\n" +
				"actual value of elem2 is different in 1 from elem4\n" +
				"actual value of elem3 is not expected in 2\n",
		},
		{
			name:     "should return expected message when maps are not equal",
			actual:   values.NewKeyStringMap(map[string]int{"1": 2, "2": 2}),
			expected: map[string]int{"1": 1},
			expectedMessage: "assertion failed:\n" +
				"expected value\t:map[1:1]\n" +
				"actual value\t:map[1:2 2:2]\n" +
				"actual value of 2 is different in 1 from 1\n" +
				"actual value of 2 is not expected in 2\n",
		},
		{
			name:            "should return expected message when simple types are not equal",
			actual:          values.NewStringValue("i'm a simple type"),
			expected:        "i'm a simple type but not equal",
			expectedMessage: "assertion failed:\nexpected value\t:i'm a simple type but not equal\nactual value\t:i'm a simple type\n",
		},
		{
			name: "should return expected message when structs are not equal",
			actual: values.NewStructValue(assertedStruct{
				SliceField: []string{"elem3"},
			}),
			expected: assertedStruct{
				BoolField:   true,
				StringField: "some-value",
				IntField:    100,
				SliceField:  []string{"elem1", "elem2"},
			},
			expectedMessage: "assertion failed:\nexpected value\t:" +
				"{BoolField:true StringField:some-value IntField:100 SliceField:[elem1 elem2]}\n" +
				"actual value\t:{BoolField:false StringField: IntField:0 SliceField:[elem3]}\n" +
				"actual value of false is different in BoolField from true\n" +
				"actual value of  is different in StringField from some-value\n" +
				"actual value of 0 is different in IntField from 100\n" +
				"actual value of elem3 is different in SliceField:0 from elem1\n" +
				"actual value of <nil> is expected but missing from SliceField:1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualMessage := shouldBeEqual(tt.actual, tt.expected)
			That(t, actualMessage).IsEqualTo(tt.expectedMessage)
		})
	}
}
