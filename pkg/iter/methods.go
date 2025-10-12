package iter

// FilterFunc is a function which takes a value and returns
// true if the value should be yielded.
type FilterFunc[T any] func(T) bool

// Filter returns an iterator which yields the elements
// from the original iterator which return true from the
// fn provided.
func (ite Iterator[T]) Filter(fn FilterFunc[T]) Iterator[T] {
	return func(yield func(T) bool) {
		for v := range ite {
			if fn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// All returns true if all the values in the iterator
// return true from the fn provided.
func (ite Iterator[T]) All(fn FilterFunc[T]) bool {
	for v := range ite {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Partition returns two slices, one for the true values
// and one for the false values.
func (ite Iterator[T]) Partition(fn FilterFunc[T]) (
	trueV []T, falseV []T) {

	trueV = make([]T, 0)
	falseV = make([]T, 0)

	for v := range ite {
		if fn(v) {
			trueV = append(trueV, v)
		} else {
			falseV = append(falseV, v)
		}
	}

	return trueV, falseV
}

// Count returns the number of elements in the iterator.
func (ite Iterator[T]) Count() int {
	count := 0
	for range ite {
		count++
	}
	return count
}

// AdvanceBy moves the iterator forward by up to n elements.
func (ite Iterator[T]) AdvanceBy(n int) []T {
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

// Nth returns the nth element in the iterator, and true if the
// value exists. Else, returns nil and false.
func (ite Iterator[T]) Nth(n int) (*T, bool) {
	if n < 0 {
		return nil, false
	}

	count := -1
	for v := range ite {
		count++
		if count == n {
			return &v, true
		}
	}

	return nil, false
}

// FirstN returns the first n elements in the iterator.
func (ite Iterator[T]) FirstN(n int) Iterator[T] {
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

// Last returns the last element in the iterator.
func (ite Iterator[T]) Last() (*T, bool) {
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

// Intersperse returns an iterator which yields the elements
// from the original iterator, with the separator between
// each element.
func (ite Iterator[T]) Intersperse(separator T) Iterator[T] {
	return func(yield func(T) bool) {
		for v := range ite {
			if !yield(v) || !yield(separator) {
				return
			}
		}
	}
}

// Map returns an iterator which yields the elements
// from the original iterator, after applying the fn
// to each element.
func Map[A any, B any](ite Iterator[A], fn func(A) B) Iterator[B] {
	return func(yield func(B) bool) {
		for v := range ite {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// Enumerate returns an iterator which yields the elements from the original
// iterator, along with the index of the element.
func (ite Iterator[T]) Enumerate() Iterator2[int, T] {
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

// Collect returns a slice containing all the elements from the iterator.
func (ite Iterator[T]) Collect() []T {
	acc := make([]T, 0)
	for v := range ite {
		acc = append(acc, v)
	}
	return acc
}

// Fold returns the result of folding the iterator using the fn provided.
func Fold[A any, B any](ite Iterator[A], acc B, fn func(A, B) B) B {
	for v := range ite {
		acc = fn(v, acc)
	}
	return acc
}

// Reduce returns the result of reducing the iterator using the fn provided.
func (ite Iterator[T]) Reduce(fn func(T, T) T) T {
	var acc *T
	for v := range ite {
		if acc == nil {
			acc = &v
			continue
		}

		*acc = fn(*acc, v)
	}

	return *acc
}

// ForEach calls the fn for each element in the iterator.
func (ite Iterator[T]) ForEach(fn func(T)) {
	for v := range ite {
		fn(v)
	}
}
