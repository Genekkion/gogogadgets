package orderedmap

import (
	"cmp"

	pqueue "github.com/Genekkion/gogogadgets/pkg/dsa/priorityqueue/simpleheap"
	"github.com/Genekkion/gogogadgets/pkg/iter"
)

type OrderedMap[K cmp.Ordered, V any] struct {
	keys   pqueue.PQueue[K]
	values map[K]V
}

func New[K cmp.Ordered, V any](priority pqueue.PriorityType, capacity ...int,
) OrderedMap[K, V] {
	var c int
	if len(capacity) > 0 {
		c = max(0, capacity[0])
	} else {
		c = 0
	}

	queue := pqueue.New[K]()

	return OrderedMap[K, V]{
		keys:   *queue,
		values: make(map[K]V, c),
	}
}

func (m *OrderedMap[K, V]) Set(k K, v V) *V {
	old, ok := m.values[k]
	if !ok {
		m.keys.Push(k)
	}
	m.values[k] = v

	if !ok {
		return nil
	}
	return &old
}

func (m *OrderedMap[K, V]) Get(k K) (*V, bool) {
	v, ok := m.values[k]
	if !ok {
		return nil, false
	}
	return &v, true
}

func (m *OrderedMap[K, V]) Keys() iter.Iterator[K] {
	return m.keys.Iter()
}

func (m *OrderedMap[K, V]) Values() iter.Iterator[V] {
	return func(yield func(V) bool) {
		for k := range m.Keys() {
			v, _ := m.Get(k)
			if !yield(*v) {
				return
			}
		}
	}
}

func (m *OrderedMap[K, V]) KV() iter.Iterator2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range m.Keys() {
			v, _ := m.Get(k)
			if !yield(k, *v) {
				return
			}
		}
	}
}
