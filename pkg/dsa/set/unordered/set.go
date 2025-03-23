package set

var emptyValue = struct{}{}

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](opts ...SetOption[T]) *Set[T] {
	s := &Set[T]{
		m: make(map[T]struct{}),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s Set[T]) Len() int {
	return len(s.m)
}

func (s Set[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

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
