package iter

import "iter"

// Iterator is a wrapper around iter.Seq[T] for convenience.
type Iterator[T any] iter.Seq[T]

// Iterator returns the iterator as iter.Seq[T].
func (ite Iterator[T]) Iterator() iter.Seq[T] {
	return iter.Seq[T](ite)
}

// Iterator2 is a wrapper around iter.Seq2[K, V] for convenience.
type Iterator2[K any, V any] iter.Seq2[K, V]

// FromSlice creates an iterator from a slice.
func FromSlice[T any](slice []T) Iterator[T] {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

// FromMap creates an iterator from a map.
func FromMap[K comparable, V any](m map[K]V) Iterator2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}
