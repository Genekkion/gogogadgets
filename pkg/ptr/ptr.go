package ptr

// New returns a pointer to the given value.
func New[T any](v T) *T {
	return &v
}
