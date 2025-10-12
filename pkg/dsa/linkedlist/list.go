package list

import (
	"github.com/Genekkion/gogogadgets/pkg/dsa/node"
)

// LinkedList is a generic doubly-linked list.
type LinkedList[T any] struct {
	head *node.LinearNode[T]
	tail *node.LinearNode[T]
	size int
}

// New creates a new linked list.
func New[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// FromExisting creates a new linked list from an existing node.
func FromExisting[T any](head *node.LinearNode[T]) LinkedList[T] {
	l := LinkedList[T]{
		head: head,
	}

	l.size, l.tail = l.countLength()

	return l
}

// countLength returns the length of the linked list and the last node.
func (l LinkedList[T]) countLength() (int, *node.LinearNode[T]) {
	if l.head == nil {
		return 0, nil
	}

	ptr := l.head
	count := 1
	for ptr.GetNext() != nil {
		ptr = ptr.GetNext()
		count++
	}
	return count, ptr
}

// Len returns the length of the linked list.
func (l LinkedList[T]) Len() int {
	return l.size
}

// IsEmpty returns true if the linked list is empty.
func (l LinkedList[T]) IsEmpty() bool {
	return l.Len() == 0
}

// PopFront removes the first node from the linked list and returns it.
func (l *LinkedList[T]) PopFront() *node.LinearNode[T] {
	if l.Len() == 0 {
		return nil
	}

	n := l.head

	l.head = l.head.GetNext()
	l.size--

	if l.size == 0 {
		l.tail = nil
	} else {
		l.head.SetPrev(nil)
	}
	return n
}

// PopBack removes the last node from the linked list and returns it.
func (l *LinkedList[T]) PopBack() *node.LinearNode[T] {
	if l.Len() == 0 {
		return nil
	}

	n := l.tail

	l.tail = l.tail.GetPrev()
	l.size--

	if l.size == 0 {
		l.head = nil
	} else {
		l.tail.SetNext(nil)
	}
	return n
}

// PushBack adds a node to the end of the linked list.
func (l *LinkedList[T]) PushBack(n *node.LinearNode[T]) {
	if l.IsEmpty() {
		l.head = n
		l.tail = n
		l.size = 1
		return
	}

	l.tail.SetNext(n)
	n.SetPrev(l.tail)
	l.tail = n
	l.size++
}

// PushFront adds a node to the beginning of the linked list.
func (l *LinkedList[T]) PushFront(n *node.LinearNode[T]) {
	if l.IsEmpty() {
		l.head = n
		l.tail = n
		l.size = 1
		return
	}

	l.head.SetPrev(n)
	n.SetNext(l.head)
	l.head = n
	l.size++
}
