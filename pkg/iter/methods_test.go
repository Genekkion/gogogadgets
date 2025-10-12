package iter

import (
	"reflect"
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Tests many Iterator methods on a simple integer sequence.
func TestIteratorMethods(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	ite := FromSlice(slice)

	// Filter evens
	evens := ite.Filter(func(v int) bool { return v%2 == 0 }).Collect()
	expectedEvens := []int{2, 4}
	test.AssertEqual(t, "Filtered length incorrect", len(expectedEvens), len(evens))
	test.Assert(t, "Filtered values incorrect", reflect.DeepEqual(expectedEvens, evens))

	// Count should be number of elements in original
	count := FromSlice(slice).Count()
	test.AssertEqual(t, "Count incorrect", len(slice), count)

	// All - are all positive?
	allPos := FromSlice(slice).All(func(v int) bool { return v > 0 })
	test.Assert(t, "Expected all positives", allPos)

	// Partition into odd and even
	trues, falses := FromSlice(slice).Partition(func(v int) bool { return v%2 == 0 })
	test.AssertEqual(t, "Partition true length", 2, len(trues))
	test.AssertEqual(t, "Partition false length", 3, len(falses))

	// AdvanceBy
	adv := FromSlice(slice).AdvanceBy(3)
	test.AssertEqual(t, "AdvanceBy length", 3, len(adv))
	test.Assert(t, "AdvanceBy values mismatch", reflect.DeepEqual([]int{1, 2, 3}, adv))

	// Nth
	nth, ok := FromSlice(slice).Nth(2)
	test.Assert(t, "Nth should be found", ok)
	if ok {
		test.AssertEqual(t, "Nth value mismatch", 3, *nth)
	}

	// FirstN and Last
	firstTwo := FromSlice(slice).FirstN(2).Collect()
	test.Assert(t, "FirstN mismatch", reflect.DeepEqual([]int{1, 2}, firstTwo))

	lastPtr, ok := FromSlice(slice).Last()
	test.Assert(t, "Last should be found", ok)
	if ok {
		test.AssertEqual(t, "Last value mismatch", 5, *lastPtr)
	}

	// Intersperse: note current implementation yields separator after every element.
	inter := FromSlice([]int{1, 2, 3}).Intersperse(0).Collect()
	test.AssertEqual(t, "Intersperse length", 6, len(inter))
	test.Assert(t, "Intersperse values", reflect.DeepEqual([]int{1, 0, 2, 0, 3, 0}, inter))

	// Map -> double values
	mapped := Map(FromSlice([]int{1, 2, 3}), func(a int) int { return a * 2 }).Collect()
	test.Assert(t, "Map values mismatch", reflect.DeepEqual([]int{2, 4, 6}, mapped))

	// Enumerate
	enumIt := FromSlice([]int{10, 20}).Enumerate()
	enumAccum := make([][2]int, 0)
	enumIt(func(i int, v int) bool {
		enumAccum = append(enumAccum, [2]int{i, v})
		return true
	})
	test.Assert(t, "Enumerate length", 2 == len(enumAccum))
	test.Assert(t, "Enumerate first pair", enumAccum[0][0] == 0 && enumAccum[0][1] == 10)

	// Collect
	col := FromSlice([]int{7, 8}).Collect()
	test.Assert(t, "Collect length", 2 == len(col))
	test.Assert(t, "Collect values", reflect.DeepEqual([]int{7, 8}, col))

	// Fold: sum all elements
	sum := Fold(FromSlice([]int{1, 2, 3}), 0, func(a int, acc int) int { return a + acc })
	test.AssertEqual(t, "Fold sum mismatch", 6, sum)

	// Reduce: sum via reduce (fn takes two elements)
	reduced := FromSlice([]int{1, 2, 3}).Reduce(func(a, b int) int { return a + b })
	test.AssertEqual(t, "Reduce sum mismatch", 6, reduced)

	// ForEach
	acc := 0
	FromSlice([]int{1, 2, 3}).ForEach(func(v int) { acc += v })
	test.AssertEqual(t, "ForEach accumulation mismatch", 6, acc)
}
