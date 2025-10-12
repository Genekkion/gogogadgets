package iter

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Test CircularSlice: use FirstN to limit the infinite circular iterator.
func TestCircularSlice(t *testing.T) {
	s := []int{10, 20}
	it := Iterator[int](CircularSlice(s)).FirstN(5)
	got := it.Collect()
	expected := []int{10, 20, 10, 20, 10}

	test.AssertEqual(t, "CircularSlice produced incorrect length", len(expected), len(got))
	for i := range expected {
		test.AssertEqual(t, "CircularSlice value mismatch at index", expected[i], got[i])
	}
}

// Test CircularSlice2: collect a few index/value pairs.
func TestCircularSlice2(t *testing.T) {
	s := []string{"x", "y"}
	seq2 := CircularSlice2(s)

	collected := make([]struct {
		I int
		V string
	}, 0)

	seq2(func(i int, v string) bool {
		collected = append(collected, struct {
			I int
			V string
		}{i, v})
		return len(collected) < 3 // stop after 3 items
	})

	// Expect three collected items
	test.AssertEqual(t, "CircularSlice2 produced incorrect number of items", 3, len(collected))

	// The first two items should be (0,"x"), (1,"y")
	test.AssertEqual(t, "First index mismatch", 0, collected[0].I)
	test.AssertEqual(t, "First value mismatch", "x", collected[0].V)
	test.AssertEqual(t, "Second index mismatch", 1, collected[1].I)
	test.AssertEqual(t, "Second value mismatch", "y", collected[1].V)
}
