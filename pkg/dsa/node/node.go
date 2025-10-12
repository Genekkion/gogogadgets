package node

import "fmt"

// Node is a generic node.
type Node[T any] struct {
	v T
}

// NewNode creates a new node.
func NewNode[T any](v T) Node[T] {
	return Node[T]{
		v: v,
	}
}

// SetValue sets the value of the node.
func (n *Node[T]) SetValue(v T) {
	n.v = v
}

// GetValue returns the value of the node.
func (n Node[T]) GetValue() T {
	return n.v
}

// String returns the string representation of the value in the node.
func (n Node[T]) String() string {
	return fmt.Sprint(n.v)
}
