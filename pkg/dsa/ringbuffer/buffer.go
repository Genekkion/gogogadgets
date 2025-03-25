package ringbuffer

import "fmt"

type Buffer[T any] struct {
	config Config
	data   []T
	start  int
	end    int
	size   int
}

func New[T any](opts ...RBOption) Buffer[T] {
	buf := Buffer[T]{
		config: defaultConfig(),
		start:  0,
		end:    0,
	}

	for _, opt := range opts {
		opt(&buf.config)
	}

	buf.data = make([]T, buf.config.bufferSize)

	return buf
}

func (b Buffer[T]) Len() int {
	return b.size
}

func (b Buffer[T]) Cap() int {
	return cap(b.data)
}

func (b *Buffer[T]) Resize() {
	data := make([]T, int(float64(len(b.data))*b.config.resizeFactor))
	if b.Len() == 0 {
		b.data = data
		return
	} else if b.end > b.start {
		copy(data, b.data[b.start:b.end+1])
	} else {
		v := b.data[b.start:]

		x := len(v)
		copy(data, v)

		v = b.data[:b.end]
		copy(data[x:], v)
		b.end = x + len(v)
	}

	b.start = 0
	b.data = data
}

func (b *Buffer[T]) Pop() (*T, bool) {
	if b.Len() == 0 {
		return nil, false
	}

	v := b.data[b.start]
	if b.Len() == 1 {
		b.start = 0
		b.end = 0
	} else {
		b.start++
	}

	b.size--
	return &v, true
}

func (b *Buffer[T]) Push(items ...T) {
	for _, v := range items {
		if b.Len() == b.Cap() {
			b.Resize()
		}

		b.data[b.end] = v
		b.end++
		b.end %= b.Cap()
		b.size++
	}
}

func (b Buffer[T]) String() string {
	if b.Len() == 0 {
		return "[]"
	} else if b.end > b.start {
		return fmt.Sprint(b.data[b.start:b.end])
	}
	return fmt.Sprintf("%v%v", b.data[b.start:], b.data[:b.end])
}
