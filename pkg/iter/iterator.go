package iter

import "iter"

type Iterator[T any] iter.Seq[T]

type Iterator2[K any, V any] iter.Seq2[K, V]

func FromSlice[T any](slice []T) Iterator[T] {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

func FromMap[K comparable, V any](m map[K]V) Iterator2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}
