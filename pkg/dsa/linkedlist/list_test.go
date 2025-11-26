package list

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/dsa/node"
	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Tests that a newly created list is empty and pops return nil.
func TestNewListEmpty(t *testing.T) {
	l := New[int]()

	test.Assert(t, "List should be empty", l.IsEmpty())
	test.Assert(t, "List should have length 0", l.Len() == 0)
	test.Assert(t, "PopFront on empty list should return nil", l.PopFront() == nil)
	test.Assert(t, "PopBack on empty list should return nil", l.PopBack() == nil)
}

// Tests that a newly created list is empty and pops return nil.
func TestPushBackAndPopFront(t *testing.T) {
	l := New[int]()

	nodes := []node.LinearNode[int]{
		node.NewLinearNode(1),
		node.NewLinearNode(2),
		node.NewLinearNode(3),
	}

	for _, n := range nodes {
		l.PushBack(&n)
	}

	test.AssertEqual(t, "Unexpected length of list", len(nodes), l.Len())

	for _, n := range nodes {
		got := l.PopFront()
		test.Assert(t, "Item should not be nil", got != nil)
		test.AssertEqual(t, "Unexpected value from node", n.GetValue(), got.GetValue())
	}

	test.Assert(t, "List should be empty", l.IsEmpty())
	test.Assert(t, "List should have length 0", l.Len() == 0)

	test.Assert(t, "PopFront on empty list should return nil", l.PopFront() == nil)
	test.Assert(t, "PopBack on empty list should return nil", l.PopBack() == nil)
}

// Tests pushing to front and popping from back yields FIFO order of original
// pushes but reversed push order.
func TestPushFrontAndPopBack(t *testing.T) {
	l := New[int]()

	nodes := []node.LinearNode[int]{
		node.NewLinearNode(1),
		node.NewLinearNode(2),
		node.NewLinearNode(3),
	}

	for _, n := range nodes {
		l.PushBack(&n)
	}

	// list: 3,2,1

	test.AssertEqual(t, "Unexpected length of list", len(nodes), l.Len())

	for i := range nodes {
		n := nodes[len(nodes)-1-i]
		got := l.PopBack()
		test.Assert(t, "Item should not be nil", got != nil)
		test.AssertEqual(t, "Unexpected value from node", n.GetValue(), got.GetValue())
	}

	test.Assert(t, "List should be empty", l.IsEmpty())
	test.Assert(t, "List should have length 0", l.Len() == 0)
}

// Tests that FromExisting correctly computes length and tail from a pre-linked
// chain.
func TestFromExisting(t *testing.T) {
	// Build chain n1 <-> n2 <-> n3
	n1 := node.NewLinearNode(10)
	n2 := node.NewLinearNode(20)
	n3 := node.NewLinearNode(30)

	// link nodes
	n1.SetNext(&n2)
	n2.SetPrev(&n1)
	n2.SetNext(&n3)
	n3.SetPrev(&n2)

	l := FromExisting[int](&n1)

	test.Assert(t, "List should not be empty", !l.IsEmpty())
	test.Assert(t, "List has incorrect length", l.Len() == 3)

	// PopBack should yield 30, then 20, then 10
	got := l.PopBack()
	test.Assert(t, "Item should not be nil", got != nil)
	test.AssertEqual(t, "Unexpected value from node", 30, got.GetValue())

	got = l.PopBack()
	test.Assert(t, "Item should not be nil", got != nil)
	test.AssertEqual(t, "Unexpected value from node", 20, got.GetValue())

	got = l.PopBack()
	test.Assert(t, "Item should not be nil", got != nil)
	test.AssertEqual(t, "Unexpected value from node", 10, got.GetValue())

	test.Assert(t, "List should be empty", l.IsEmpty())
	test.Assert(t, "List should have length 0", l.Len() == 0)
}
