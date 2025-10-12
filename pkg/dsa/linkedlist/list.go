package list

import (
	"fmt"
	"strings"

	"github.com/Genekkion/gogogadgets/pkg/dsa/node"
)

type LinkedList[T any] struct {
	head *node.LinearNode[T]
	tail *node.LinearNode[T]
	size int
}

func New[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func FromExisting[T any](head *node.LinearNode[T]) LinkedList[T] {

	ptr := head
	count := 0
	for ptr != nil {
		count++
		ptr = ptr.GetNext()
	}

	return LinkedList[T]{
		head: head,
		tail: ptr,
		size: count,
	}
}

func (l LinkedList[T]) Len() int {
	return l.size
}

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

func (l *LinkedList[T]) PushBack(n *node.LinearNode[T]) {
	l.tail.SetNext(n)
	n.SetPrev(l.tail)
	l.tail = n
	l.size++
}

func (l *LinkedList[T]) PushFront(n *node.LinearNode[T]) {
	l.head.SetPrev(n)
	n.SetNext(l.head)
	l.head = n
	l.size++
}

func (l LinkedList[T]) String() string {
	b := strings.Builder{}
	b.WriteRune('[')
	ptr := l.head
	if ptr != nil {
		b.WriteString(fmt.Sprint(ptr.GetValue()))
		ptr = ptr.GetNext()

		for ptr != nil {
			b.WriteString(", ")
			b.WriteString(fmt.Sprint(ptr.GetValue()))
			ptr = ptr.GetNext()
		}
	}
	b.WriteRune(']')

	return b.String()
}
