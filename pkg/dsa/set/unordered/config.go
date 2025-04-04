package set

type SetOption[T comparable] func(*Set[T])

func WithSlice[T comparable](slice []T) SetOption[T] {
	return func(s *Set[T]) {
		for _, v := range slice {
			s.m[v] = emptyValue
		}
	}
}
