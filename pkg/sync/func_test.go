package sync

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

// Tests CallOnceAfter in a single-threaded context: fn should be called exactly
// once.
func TestCallOnceAfterSingleThread(t *testing.T) {
	calls := atomic.Int32{}
	calls.Store(0)
	fn := func() {
		calls.Add(1)
	}

	f := CallOnceAfter(3, fn)

	// Call 5 times; fn should run exactly once (on the 3rd call).
	for i := 0; i < 5; i++ {
		f()
	}

	got := calls.Load()

	test.AssertEqual(t, "Expected a single call", 1, got)
}

// Tests CallOnceAfter under concurrent calls: fn should still be called exactly
// once.
func TestCallOnceAfterConcurrent(t *testing.T) {
	calls := atomic.Int32{}
	calls.Store(0)
	fn := func() {
		calls.Add(1)
	}

	n := int32(10)
	f := CallOnceAfter(n, fn)

	total := 20
	wg := sync.WaitGroup{}
	wg.Add(total)

	for range total {
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()

	got := calls.Load()
	test.AssertEqual(t, "Expected different number of calls", 1, int(got))
}

// Tests CallAfter in a single-threaded context: fn should be called for the nth
// and later calls.
func TestCallAfterSingleThread(t *testing.T) {
	calls := atomic.Int32{}
	calls.Store(0)
	fn := func() {
		calls.Add(1)
	}

	n := int32(3)
	f := CallAfter(n, fn)

	// Call 5 times; fn should run on calls 3,4,5 => 3 times.
	total := 5
	for range total {
		f()
	}

	got := calls.Load()
	test.AssertEqual(t, "Unexpected number of calls", total-int(n)+1, int(got))
}

// Tests CallAfter under concurrent calls: fn should be called totalCalls - n + 1
// times when totalCalls >= n.
func TestCallAfterConcurrent(t *testing.T) {
	calls := atomic.Int32{}
	calls.Store(0)
	fn := func() {
		calls.Add(1)
	}

	n := int32(10)
	f := CallAfter(n, fn)

	total := 20
	wg := sync.WaitGroup{}
	wg.Add(total)

	for range total {
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()

	got := calls.Load()
	test.AssertEqual(t, "Unexpected number of calls", total-int(n)+1, int(got))
}
