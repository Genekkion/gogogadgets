package btree

import "github.com/Genekkion/gogogadgets/pkg/dsa/node"

type TraverseFunc[T any] func(T)

func (t BTree[T]) DFS(traverseFuncs ...TraverseFunc[T]) {
	var recAux func(*node.BTreeNode[T])
	recAux = func(root *node.BTreeNode[T]) {
		if root == nil {
			return
		}

		v := root.GetValue()
		for _, f := range traverseFuncs {
			f(v)
		}

		recAux(root.GetLeft())
		recAux(root.GetRight())
	}

	recAux(t.head)
}

func (t BTree[T]) Preorder(traverseFuncs ...TraverseFunc[T]) {
	t.DFS(traverseFuncs...)
}

func (t BTree[T]) Inorder(traverseFuncs ...TraverseFunc[T]) {
	var recAux func(*node.BTreeNode[T])
	recAux = func(root *node.BTreeNode[T]) {
		if root == nil {
			return
		}

		recAux(root.GetLeft())

		v := root.GetValue()
		for _, f := range traverseFuncs {
			f(v)
		}

		recAux(root.GetRight())
	}

	recAux(t.head)
}

func (t BTree[T]) Postorder(traverseFuncs ...TraverseFunc[T]) {
	var recAux func(*node.BTreeNode[T])
	recAux = func(root *node.BTreeNode[T]) {
		if root == nil {
			return
		}

		recAux(root.GetLeft())
		recAux(root.GetRight())

		v := root.GetValue()
		for _, f := range traverseFuncs {
			f(v)
		}
	}

	recAux(t.head)
}

func (t BTree[T]) BFS(traverseFuncs ...TraverseFunc[T]) {
	queue := []*node.BTreeNode[T]{t.head}

	for len(queue) != 0 {
		n := queue[0]

		v := n.GetValue()
		for _, f := range traverseFuncs {
			f(v)
		}

		queue = queue[1:]

		l := n.GetLeft()
		r := n.GetRight()
		if l != nil {
			queue = append(queue, l)
		}
		if r != nil {
			queue = append(queue, r)
		}
	}
}
