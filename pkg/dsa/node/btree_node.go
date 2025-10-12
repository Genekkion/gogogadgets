package node

// BTreeNode Type for a generic binary tree node.
type BTreeNode[T any] struct {
	Node[T]

	left  *BTreeNode[T]
	right *BTreeNode[T]
}

// NewBTreeNode creates a new binary tree node.
func NewBTreeNode[T any](v T) BTreeNode[T] {
	return BTreeNode[T]{
		Node:  NewNode(v),
		left:  nil,
		right: nil,
	}
}

// GetLeft returns the left child.
func (n BTreeNode[T]) GetLeft() *BTreeNode[T] {
	return n.left
}

// SetLeft sets the left child.
func (n *BTreeNode[T]) SetLeft(left *BTreeNode[T]) {
	n.left = left
}

// GetRight returns the right child.
func (n BTreeNode[T]) GetRight() *BTreeNode[T] {
	return n.right
}

// SetRight sets the right child.
func (n *BTreeNode[T]) SetRight(right *BTreeNode[T]) {
	n.right = right
}
