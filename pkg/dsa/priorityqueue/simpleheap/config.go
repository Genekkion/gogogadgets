package pqueue

import "cmp"

type PQueueOption[T cmp.Ordered] func(*PQueue[T])

func WithPriority[T cmp.Ordered](priority PriorityType) PQueueOption[T] {
	return func(p *PQueue[T]) {
		p.priority = priority
	}
}

func WithCapacity[T cmp.Ordered](capacity int) PQueueOption[T] {
	capacity = max(0, capacity)

	return func(p *PQueue[T]) {
		if len(p.data) < capacity {
			p.data = make([]T, 0, capacity)
		}
	}
}

func WithItems[T cmp.Ordered](items ...T) PQueueOption[T] {
	return func(p *PQueue[T]) {
		p.data = append(p.data, items...)
	}
}
