package node

import "fmt"

type Node[T any] struct {
	v T
}

func New[T any](v T) Node[T] {
	return Node[T]{
		v: v,
	}
}

func (n *Node[T]) SetValue(v T) {
	n.v = v
}

func (n Node[T]) GetValue() T {
	return n.v
}

func (n Node[T]) String() string {
	return fmt.Sprint(n.v)
}
