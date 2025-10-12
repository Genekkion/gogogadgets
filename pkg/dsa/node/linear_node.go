package node

// LinearNode is a node meant for linear data structures such as linked lists.
type LinearNode[T any] struct {
	Node[T]

	prev *LinearNode[T]
	next *LinearNode[T]
}

// NewLinearNode creates a new linear node.
func NewLinearNode[T any](v T) LinearNode[T] {
	return LinearNode[T]{
		Node: NewNode(v),
		prev: nil,
		next: nil,
	}
}

// GetPrev returns the previous node.
func (n LinearNode[T]) GetPrev() *LinearNode[T] {
	return n.prev
}

// SetPrev sets the previous node.
func (n *LinearNode[T]) SetPrev(prev *LinearNode[T]) {
	n.prev = prev
}

// GetNext returns the next node.
func (n LinearNode[T]) GetNext() *LinearNode[T] {
	return n.next
}

// SetNext sets the next node.
func (n *LinearNode[T]) SetNext(next *LinearNode[T]) {
	n.next = next
}
