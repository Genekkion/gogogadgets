package iter

import (
	"reflect"
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Test FromSlice creates an iterator from a slice and collects the values.
func TestFromSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	it := FromSlice(slice)

	collected := it.Collect()
	test.AssertEqual(t, "Collected length mismatch", len(slice), len(collected))
	test.Assert(t, "Collected values mismatch", reflect.DeepEqual(slice, collected))
}

// Test FromMap creates an iterator from a map and iterates over key/value pairs.
func TestFromMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	it := FromMap(m)

	got := map[string]int{}
	it(func(k string, v int) bool {
		got[k] = v
		return true
	})

	test.AssertEqual(t, "Map length mismatch", len(m), len(got))

	for k, v := range m {
		val, ok := got[k]
		test.Assert(t, "Expected key missing from iterator results", ok)
		test.AssertEqual(t, "Value mismatch for key "+k, v, val)
	}
}
