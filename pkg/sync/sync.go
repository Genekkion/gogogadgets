package sync

import "sync/atomic"

// fn is only called exactly once after n calls to the
// function returned. Thread safe.
func CallOnceAfter(n int32, fn func()) func() {
	counter := atomic.Int32{}
	counter.Store(0)

	return func() {
		if counter.Add(1) == n {
			fn()
		}
	}
}

// fn is only called after n calls to the function returned.
// Inclusive of the nth call. Thread safe.
func CallAfter(n int32, fn func()) func() {
	counter := atomic.Int32{}
	counter.Store(0)

	return func() {

		if counter.Add(1) >= n {
			fn()
		}
	}
}
