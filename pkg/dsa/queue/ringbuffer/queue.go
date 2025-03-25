package queue

import "github.com/Genekkion/gogogadgets/pkg/dsa/ringbuffer"

type Queue[T any] ringbuffer.Buffer[T]

func New[T any](opts ...ringbuffer.RBOption) Queue[T] {
	return Queue[T](ringbuffer.New[T](opts...))
}
