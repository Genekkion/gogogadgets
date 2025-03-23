package result

type Result[T any] struct {
	v   *T
	err error
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		v:   nil,
		err: err,
	}
}

func Ok[T any](v T) Result[T] {
	return Result[T]{
		v:   &v,
		err: nil,
	}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

type MapOkFunc[A any, B any] func(A) B
type MapErrFunc func(error) error

// Ignores the map function if result is already an error
func MapOk[A any, B any](r *Result[A], mapFunc MapOkFunc[A, B]) Result[B] {
	if r.IsErr() {
		return Err[B](r.err)
	}

	return Ok(mapFunc(*(r.v)))
}

// Ignores the map function if result is ok
func MapErr[A any](r *Result[A], mapFunc MapErrFunc) Result[A] {
	if r.IsOk() {
		return *r
	}

	return Err[A](mapFunc(r.err))
}

// Will panic if its an error
func (r Result[T]) unwrap(msg ...string) T {
	if r.IsErr() {
		if len(msg) > 0 {
			panic(msg[0])
		} else {
			panic("result is an error")
		}
	}
	return *(r.v)
}

// Will panic if its an error
func (r Result[T]) Expect(msg string) T {
	return r.unwrap(msg)
}

// Will panic if its an error
func (r Result[T]) Unwrap() T {
	return r.unwrap()
}

// Will panic if its an error
func (r Result[T]) UnwrapOr(v T) T {
	if r.IsErr() {
		return v
	}
	return *(r.v)
}
