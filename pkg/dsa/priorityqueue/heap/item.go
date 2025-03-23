package pqueue

type Item interface {
	Less(another Item) bool
}
