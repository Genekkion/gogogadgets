package pqueue

type PQueueOption[T Item] func(*PQueue[T])

func WithPriority[T Item](priority PriorityType) PQueueOption[T] {
	return func(p *PQueue[T]) {
		p.priority = priority
	}
}

func WithCapacity[T Item](capacity int) PQueueOption[T] {
	capacity = max(0, capacity)

	return func(p *PQueue[T]) {
		if len(p.data) < capacity {
			p.data = make([]T, 0, capacity)
		}
	}
}

func WithItems[T Item](items ...T) PQueueOption[T] {
	return func(p *PQueue[T]) {
		p.data = append(p.data, items...)
	}
}
