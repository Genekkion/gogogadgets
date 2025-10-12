package btree

import (
	"fmt"

	"github.com/Genekkion/gogogadgets/pkg/dsa/node"
)

type BTree[T any] struct {
	head *node.BTreeNode[T]
	size int
}

func New[T any](head *node.BTreeNode[T]) BTree[T] {
	t := BTree[T]{
		head: head,
		size: 0,
	}

	t.DFS(func(T) {
		t.size++
	})

	return t
}

func (t BTree[T]) String() string {
	return fmt.Sprintf("BTree: { head: %v, size: %d }",
		t.head.GetValue(), t.size)
}
