package iter

import "iter"

// Zip merges two iterators into a single iterator.
func Zip[A any, B any](ite1 Iterator[A], ite2 Iterator[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		stopped1 := false
		stopped2 := false

		next1, stop1 := iter.Pull(ite1.Iterator())
		next2, stop2 := iter.Pull(ite2.Iterator())

		for {
			v1, ok1 := next1()
			if !ok1 {
				stop1()
				stopped1 = true
			}
			v2, ok2 := next2()
			if !ok2 {
				stop2()
				stopped2 = true
			}

			if stopped1 && stopped2 {
				return
			} else if stopped1 {
				stop2()
				return
			} else if stopped2 {
				stop1()
				return
			} else if !yield(v1, v2) {
				return
			}
		}
	}
}
