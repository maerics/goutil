package goutil

import (
	"sync"
	"time"

	"github.com/maerics/golog"
)

func DoBatchWork[T any](size int, items []T, f func(int, T)) {
	if size < 1 {
		size = 1
	}

	golog.Printf("starting parallel work for %v item(s) with batch size %v", len(items), size)
	t0 := time.Now()
	wg := &sync.WaitGroup{}
	// Spawn the workers.
	for i := 0; i < size; i++ {
		wg.Add(1)
		time.Sleep(time.Duration(i) * time.Millisecond) // So they start in order.
		// Spawn a goroutine to do all the work with index as a factor of its own.
		go func(threadId int) {
			debugf := debugOrNoop(threadId < len(items))
			debugf("spawning worker #%v", threadId)
			for currentIdx := threadId; currentIdx < len(items); currentIdx += size {
				debugf("worker #%v executing item #%v (%#v)", threadId, currentIdx, items[currentIdx])
				f(currentIdx, items[currentIdx])
			}
			debugf("worker #%v finished", threadId)
			wg.Done()
		}(i)
	}
	wg.Wait()
	golog.Debugf("parallel work of %v item(s) completed in %v", len(items), time.Since(t0))
}

func debugOrNoop(b bool) func(string, ...any) {
	if b {
		return golog.Debugf
	}
	return func(s string, a ...any) {}
}
