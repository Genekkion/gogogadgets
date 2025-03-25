package iter

import (
	"context"
	"github.com/Genekkion/gogogadgets/pkg/sync"
	"iter"
	"runtime"
)

var (
	nCpu = max(1, runtime.NumCPU()-1)

	bufferSize = 64
)

// Does not preserve order
func ParallelMap[A any, B any](ite Iterator[A], fn func(A) B,
	maxCpuCount ...int) iter.Seq[B] {

	return func(yield func(B) bool) {
		var nC int
		if len(maxCpuCount) > 0 {
			nC = max(1, maxCpuCount[0])
		} else {
			nC = nCpu
		}

		in := make(chan A, nC)
		out := make(chan B, nC*bufferSize)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		closeOut := sync.CallOnceAfter(nC, func() {
			close(out)
		})

		for range nC {
			go func() {
				defer closeOut()

				for v := range in {
					select {
					case <-ctx.Done():
						return
					default:
					}

					x := fn(v)

					select {
					case out <- x:
						// TODO:
					case <-ctx.Done():
						return
					}
				}
			}()
		}

		go func() {
			defer close(in)

			for v := range ite {
				select {
				case in <- v:
				case <-ctx.Done():
					return
				}
			}
		}()

		for v := range out {
			if !yield(v) {
				cancel()
				return
			}
		}
	}
}

// Does not preserve order
func ParallelFilter[T any](ite iter.Seq[T], filterFunc FilterFunc[T],
	maxCpuCount ...int) iter.Seq[T] {

	return func(yield func(T) bool) {
		var nC int
		if len(maxCpuCount) > 0 {
			nC = max(1, maxCpuCount[0])
		} else {
			nC = nCpu
		}

		in := make(chan T, nC)
		out := make(chan T, nC*bufferSize)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		closeOut := sync.CallOnceAfter(nC, func() {
			close(out)
		})

		for range nC {
			go func() {
				defer closeOut()

				for v := range in {
					select {
					case <-ctx.Done():
						return
					default:
					}

					if filterFunc(v) {
						select {
						case out <- v:
							// TODO:
							// charm.Debug("Handled", "i", i)
						case <-ctx.Done():
							return
						}
					}
				}
			}()
		}

		go func() {
			defer close(in)

			for v := range ite {
				select {
				case in <- v:
				case <-ctx.Done():
					return
				}
			}
		}()

		for v := range out {
			if !yield(v) {
				cancel()
				return
			}
		}
	}
}
