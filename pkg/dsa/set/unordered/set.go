package set

// For use with the underlying map value.
var emptyValue = struct{}{}

// The generic set type.
type Set[T comparable] struct {
	m map[T]struct{}
}

// Creates a new set, with the options specified.
func New[T comparable](opts ...SetOption[T]) Set[T] {
	s := Set[T]{
		m: make(map[T]struct{}),
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// Returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s.m)
}

// Returns whether the set contains the specified key.
func (s Set[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Adds the keys to the set. Returns whether the set has
// been modified.
func (s *Set[T]) Add(keys ...T) bool {
	return s.Set(keys...)
}

// Read Add(keys ...T) for the description.
func (s *Set[T]) Set(keys ...T) (modified bool) {
	modified = false

	for _, key := range keys {
		v := s.Contains(key)
		if v {
			continue
		}

		s.m[key] = emptyValue
		modified = true
	}

	return modified
}

// Removes the keys specified from the set. Returns whether
// the set has been modified.
func (s *Set[T]) Remove(keys ...T) (modified bool) {
	modified = false

	for _, key := range keys {
		v := s.Contains(key)
		if !v {
			continue
		}

		delete(s.m, key)
		modified = true
	}
	return modified
}
