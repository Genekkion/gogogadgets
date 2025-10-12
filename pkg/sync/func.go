package sync

import "sync/atomic"

// CallOnceAfter calls the function fn exactly once on the nth call. Thread safe.
func CallOnceAfter(n int32, fn func()) func() {
	counter := atomic.Int32{}
	counter.Store(0)

	return func() {
		if counter.Add(1) == n {
			fn()
		}
	}
}

// CallAfter calls the function fn after n calls. Thread safe.
func CallAfter(n int32, fn func()) func() {
	counter := atomic.Int32{}
	counter.Store(0)

	return func() {
		if counter.Add(1) >= n {
			fn()
		}
	}
}
