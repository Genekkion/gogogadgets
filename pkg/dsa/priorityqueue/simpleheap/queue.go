package pqueue

import (
	"cmp"
	"container/heap"

	"github.com/Genekkion/gogogadgets/pkg/iter"
	log "github.com/Genekkion/gogogadgets/pkg/log/charm"
)

type PriorityType int

const (
	PriorityMin PriorityType = iota
	PriorityMax
)

const (
	defaultPriority = PriorityMax
)

type PQueue[T cmp.Ordered] struct {
	priority PriorityType
	data     []T
}

// Defaults to a queue which prioritises a higher priority
func New[T cmp.Ordered](opts ...PQueueOption[T]) *PQueue[T] {
	q := &PQueue[T]{
		priority: defaultPriority,
	}

	for _, opt := range opts {
		opt(q)
	}

	heap.Init(q)
	return q
}

func (q PQueue[T]) Len() int {
	return len(q.data)
}

func (q PQueue[T]) Less(i int, j int) bool {
	v := cmp.Less(q.data[i], q.data[j])

	switch q.priority {
	case PriorityMax:
		return !v
	case PriorityMin:
		return v
	default:
		panic("should not reach here")
	}
}

func (q PQueue[T]) Swap(i int, j int) {
	if q.Len() <= 0 {
		return
	}
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *PQueue[T]) Push(v any) {
	ele, ok := v.(T)
	if !ok {
		log.Error("type mismatch, v does not have a compatible type, skipping push operation")
		return
	}
	(*q).data = append((*q).data, ele)
}

// Will return nil if empty
func (q *PQueue[T]) Pop() any {
	n := len((*q).data)

	if n == 0 {
		return nil
	}

	n--
	v := (*q).data[n]
	(*q).data = (*q).data[:n]

	return v
}

// Creates an immutable clone of the original
func (q *PQueue[T]) IterClone() iter.Iterator[T] {
	clone := PQueue[T]{
		data:     make([]T, q.Len()),
		priority: q.priority,
	}
	copy(clone.data, (*q).data)

	return func(yield func(T) bool) {
		for clone.Len() > 0 {
			v := heap.Pop(&clone)
			if v == nil {
				return
			} else if !yield(v.(T)) {
				return
			}
		}
	}
}

func (q *PQueue[T]) Iter() iter.Iterator[T] {
	return func(yield func(T) bool) {
		for q.Len() > 0 {
			v := heap.Pop(q)
			if v == nil {
				return
			} else if !yield(v.(T)) {
				return
			}
		}
	}
}
