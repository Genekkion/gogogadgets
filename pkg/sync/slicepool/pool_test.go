package slicepool

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Tests that Get() returns a slice with the correct length and capacity.
func TestPoolGetReturnsMinLenAndCap(t *testing.T) {
	minSize := 2
	maxSize := 5
	p := New[int](minSize, maxSize)

	buf := p.Get()

	test.AssertEqual(t, "Incorrect length of slice returned", minSize, len(buf))
	test.AssertEqual(t, "Incorrect capacity of slice returned", maxSize, cap(buf))
}

// Tests that Put() resets the slice's length to the max size.
func TestPoolPutResetsLenToMax(t *testing.T) {
	minSize := 2
	maxSize := 5
	p := New[int](minSize, maxSize)

	// Get a buffer and put it back as-is.
	buf := p.Get()
	p.Put(buf)

	buf = p.Get()
	test.AssertEqual(t, "Incorrect length of slice returned", maxSize, len(buf))
	test.AssertEqual(t, "Incorrect capacity of slice returned", maxSize, cap(buf))
}

// Tests that Put() expands the slice to the max size if the slice is shorter than the min size.
func TestPoolPutShorterSliceIsExpandedToMax(t *testing.T) {
	minSize := 3
	maxSize := 7
	p := New[int](minSize, maxSize)

	buf := p.Get()

	// Shorten the slice below minSize and put it back
	short := buf[:1]
	p.Put(short)

	got := p.Get()
	test.AssertEqual(t, "Incorrect length of slice returned", maxSize, len(got))
	test.AssertEqual(t, "Incorrect capacity of slice returned", maxSize, cap(got))
}
