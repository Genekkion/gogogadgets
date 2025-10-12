package iter

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Test Zip stops when one iterator ends.
func TestZip(t *testing.T) {
	ints := []int{1, 2, 3}
	runes := []rune{'a', 'b'}

	a := FromSlice(ints)
	b := FromSlice(runes)

	ite := Zip[int, rune](a, b)

	i := 0
	for x, y := range ite {
		v1 := ints[i]
		v2 := runes[i]
		test.AssertEqual(t, "Value mismatch", v1, x)
		test.AssertEqual(t, "Value mismatch", v2, y)
		i++
	}
}
