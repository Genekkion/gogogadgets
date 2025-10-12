package slicepool

import "sync"

// slicepool is intended to be used for reducing allocations
// by reusing slices.

// Pool is a type safe pool which stores slices of type T.
type Pool[T any] struct {
	pool    sync.Pool
	minSize int
	maxSize int
}

// New creates a new pool.
func New[T any](minSize int, maxSize int) Pool[T] {
	return Pool[T]{
		pool: sync.Pool{
			New: func() any {
				return make([]T, minSize, maxSize)
			},
		},
		minSize: minSize,
		maxSize: maxSize,
	}
}

// Get gets a slice from the pool.
func (p *Pool[T]) Get() []T {
	buf, ok := p.pool.Get().([]T)
	if !ok {
		panic("should not reach here")
	}
	return buf
}

// Put puts a slice back to the pool.
func (p *Pool[T]) Put(buf []T) {
	buf = buf[:p.maxSize]
	p.pool.Put(buf)
}
