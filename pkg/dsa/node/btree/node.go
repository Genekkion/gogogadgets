package node

import (
	"fmt"

	"github.com/Genekkion/gogogadgets/pkg/dsa/node"
)

type BinaryNode[T any] struct {
	node.Node[T]

	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func New[T any](v T) BinaryNode[T] {
	return BinaryNode[T]{
		Node:  node.New(v),
		left:  nil,
		right: nil,
	}
}

func (n BinaryNode[T]) GetLeft() *BinaryNode[T] {
	return n.left
}

func (n *BinaryNode[T]) SetLeft(left *BinaryNode[T]) {
	n.left = left
}

func (n BinaryNode[T]) GetRight() *BinaryNode[T] {
	return n.right
}

func (n *BinaryNode[T]) SetRight(right *BinaryNode[T]) {
	n.right = right
}

func (n BinaryNode[T]) String() string {
	var left string
	if n.left != nil {
		left = fmt.Sprint(n.left.Node.GetValue())
	} else {
		left = "nil"
	}

	var right string
	if n.right != nil {
		right = fmt.Sprint(n.right.Node.GetValue())
	} else {
		right = "nil"
	}

	return fmt.Sprintf("Node: { v: %v, l: %s, r: %s }", n.v, left, right)
}
