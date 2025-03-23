package iter

import (
	"iter"
)

type FilterFunc[T any] func(T) bool

func Filter[T any](ite iter.Seq[T], filterFunc FilterFunc[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range ite {
			if filterFunc(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
func All[T any](ite iter.Seq[T], filterFunc FilterFunc[T]) bool {
	for v := range ite {
		if !filterFunc(v) {
			return false
		}
	}
	return true
}

func Partition[T any](ite iter.Seq[T], filterFunc FilterFunc[T]) ([]T, []T) {
	acc1 := make([]T, 0)
	acc2 := make([]T, 0)

	for v := range ite {
		if filterFunc(v) {
			acc1 = append(acc1, v)
		} else {
			acc2 = append(acc2, v)
		}
	}

	return acc1, acc2
}

func Count[T any](ite iter.Seq[T]) int {
	count := 0
	for range ite {
		count++
	}
	return count
}

func AdvanceBy[T any](ite iter.Seq[T], n int) []T {
	if n <= 0 {
		return nil
	}

	count := 0
	acc := make([]T, 0, n)
	for v := range ite {
		acc = append(acc, v)
		count++

		if count == n {
			return acc
		}
	}

	return acc
}

func Nth[T any](ite iter.Seq[T], n int) (*T, bool) {
	if n < 0 {
		return nil, false
	}

	count := 0
	for v := range ite {
		count++
		if count == n {
			return &v, true
		}
	}

	return nil, false
}

func FirstN[T any](ite iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n < 0 {
			return
		}

		count := 0
		for v := range ite {
			count++
			if !yield(v) || count == n {
				return
			}
		}
	}
}

func Last[T any](ite iter.Seq[T]) (*T, bool) {
	var res T
	flag := false
	for v := range ite {
		res = v
		flag = true
	}
	if !flag {
		return nil, false
	}

	return &res, true
}

func Intersperse[T any](ite iter.Seq[T], seperator T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range ite {
			if !yield(v) || !yield(seperator) {
				return
			}
		}
	}
}

type MapFunc[A any, B any] func(A) B

func Map[A any, B any](ite iter.Seq[A], mapFunc MapFunc[A, B]) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range ite {
			if !yield(mapFunc(v)) {
				return
			}
		}
	}
}

func Enumerate[T any](ite iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for v := range ite {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

func Collect[T any](ite iter.Seq[T]) []T {
	acc := make([]T, 0)
	for v := range ite {
		acc = append(acc, v)
	}
	return acc
}

type FoldFunc[A any, B any] func(A, B) B

func Fold[A any, B any](ite iter.Seq[A], acc B, foldFunc FoldFunc[A, B]) B {
	for v := range ite {
		acc = foldFunc(v, acc)
	}

	return acc
}

type ReduceFunc[T any] func(T, T) T

func Reduce[T any](ite iter.Seq[T], reduceFunc ReduceFunc[T]) T {
	var acc *T
	for v := range ite {
		if acc == nil {
			acc = &v
			continue
		}

		*acc = reduceFunc(*acc, v)
	}

	return *acc
}

type EachFunc[T any] func(T)

func ForEach[T any](ite iter.Seq[T], eachFunc EachFunc[T]) {
	for v := range ite {
		eachFunc(v)
	}
}
