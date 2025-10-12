package iter

import "iter"

// RangeN returns an iterator that yields the numbers from 0 to n-1.
func RangeN(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				return
			}
		}
	}
}

// CircularSlice returns an iterator that yields the elements of the slice
// in a circular fashion.
func CircularSlice[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		n := len(slice)
		if n == 0 {
			return
		}

		i := 0
		for {
			if !yield(slice[i]) {
				return
			}

			i++
			i %= n
		}
	}
}

// CircularSlice2 returns an iterator that yields the elements of the slice
// in a circular fashion.
func CircularSlice2[T any](slice []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		n := len(slice)
		if n == 0 {
			return
		}

		i := 0
		for {
			if !yield(i, slice[i]) {
				return
			}

			i++
			i %= n
		}
	}
}
