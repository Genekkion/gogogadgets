package queue

import "github.com/Genekkion/gogogadgets/pkg/dsa/ringbuffer"

type Queue[T any] ringbuffer.Buffer[T]
