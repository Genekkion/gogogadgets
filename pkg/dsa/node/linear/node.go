package node

import "github.com/Genekkion/gogogadgets/pkg/dsa/node"

type LinearNode[T any] struct {
	node.Node[T]

	prev *LinearNode[T]
	next *LinearNode[T]
}

func New[T any](v T) LinearNode[T] {
	return LinearNode[T]{
		Node: node.New(v),
		prev: nil,
		next: nil,
	}
}

func (n LinearNode[T]) GetPrev() *LinearNode[T] {
	return n.prev
}

func (n *LinearNode[T]) SetPrev(prev *LinearNode[T]) {
	n.prev = prev
}

func (n LinearNode[T]) GetNext() *LinearNode[T] {
	return n.next
}

func (n *LinearNode[T]) SetNext(next *LinearNode[T]) {
	n.next = next
}
