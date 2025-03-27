package set

// Type for setting a config option for the set to be
// created
type SetOption[T comparable] func(*Set[T])

// Adds elements from an initial slice of elements
func WithSlice[T comparable](slice []T) SetOption[T] {
	return func(s *Set[T]) {
		for _, v := range slice {
			s.m[v] = emptyValue
		}
	}
}
