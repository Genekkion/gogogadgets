package set

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Tests creating a new set.
func TestNewSet(t *testing.T) {
	s := New[int]()
	expected := 0
	got := s.Len()

	test.AssertEqual(t, "Expected empty set", expected, got)
}

// Tests creating a new set with a slice.
func TestNewSetWithSlice(t *testing.T) {
	slice := []int{1, 2, 3, 3}

	// Includes duplicate
	s := New(WithSlice(slice))

	l := s.Len()
	test.AssertEqual(t, "Incorrect length", 3, l)

	if s.Contains(1) && s.Contains(2) && !s.Contains(3) {
		return
	}

	expected := []int{1, 2, 3}
	got := make([]int, 0, len(expected))
	for k := range s.m {
		got = append(got, k)
	}
	test.AssertEqual(t, "Unexpected set elements", expected, got)
}

// Tests adding elements to the set.
func TestAdd(t *testing.T) {
	s := New[int]()
	keys := []int{1, 2, 3}

	modified := s.Add(keys...)
	test.AssertEqual(t, "Expected set to be modified", true, modified)
	test.AssertEqual(t, "Incorrect set length", len(keys), s.Len())

	// Add duplicate elements
	modified = s.Add(2, 3)
	test.AssertEqual(t, "Expected set to not be modified", false, modified)
	test.AssertEqual(t, "Incorrect set length", len(keys), s.Len())
}

// Tests checking if a set contains an element.
func TestContains(t *testing.T) {
	s := New[int]()
	keys := []int{5, 10}
	s.Add(keys...)

	for _, k := range keys {
		test.Assert(t, "Expected set to contain value", s.Contains(k))
	}

	test.AssertEqual(t, "Expected set not to contain 7", false, s.Contains(7))
}

func TestRemove(t *testing.T) {
	s := New[int]()
	keys := []int{1, 2, 3}
	s.Add(keys...)

	// Remove existing elements
	modified := s.Remove(2)
	test.AssertEqual(t, "Expected modification when removing existing element", true, modified)
	test.AssertEqual(t, "Expected set not to contain 2 after removal", false, s.Contains(2))
	test.AssertEqual(t, "Unexpected length", len(keys)-1, s.Len())

	// Remove non-existing element
	modified = s.Remove(42)
	test.AssertEqual(t, "Expected no modification when removing non-existing element", false, modified)
	test.AssertEqual(t, "Unexpected length", len(keys)-1, s.Len())
}
