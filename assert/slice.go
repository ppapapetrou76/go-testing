package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/core"
)

// ThatSlice returns a proper assertable structure based on the slice type
func ThatSlice(t *testing.T, actual interface{}) core.AssertableSlice {
	switch v := actual.(type) {
	case []string:
		return AssertableStringSlice{
			T:      t,
			Actual: v,
		}
	default:
		return nil
	}
}
