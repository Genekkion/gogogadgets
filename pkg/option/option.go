package option

type Option[T any] struct {
	v        *T
	hasValue bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		v:        &v,
		hasValue: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		v:        nil,
		hasValue: false,
	}
}

func (o Option[T]) IsSome() bool {
	return o.hasValue
}

func (o Option[T]) IsNone() bool {
	return !o.hasValue
}

// Panics if its none
func (o Option[T]) unwrap(msg ...string) T {
	if o.IsNone() {
		if len(msg) > 0 {
			panic(msg[0])
		} else {
			panic("option is an error")
		}
	}
	return *o.v
}

// Attempts to retrieve the specified value. Panics if it is
// none. Accepts an optional message parameter for the panic
// message.
func (o Option[T]) Unwrap(msg ...string) T {
	return o.unwrap(msg...)
}

// Returns the specified value if it is some, otherwise
// the default value provided.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return *(o.v)
}
