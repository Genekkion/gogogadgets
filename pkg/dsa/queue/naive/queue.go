package queue

type Queue[T any] []T

func (q Queue[T]) Len() int {
	return len(q)
}

func (q *Queue[T]) Push(items ...T) {
	*q = append(*q, items...)
}

func (q *Queue[T]) Pop() (*T, bool) {
	if q.Len() == 0 {
		return nil, false
	}

	v := (*q)[0]
	*q = (*q)[1:]
	return &v, true
}
