package queue

type Queue[T any] interface {
	Len() int
	Pop() (*T, bool)
	Push(...T)
}
