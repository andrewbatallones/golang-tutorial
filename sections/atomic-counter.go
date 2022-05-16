package sections

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounters() {
	fmt.Println("=== Regular Counter ===")
	regularCounters()

	fmt.Println("=== Atomic Counter ===")
	atomicCounters()

	// You can see that because there are multiple goroutines going, that it will overwrite that variable
	// With atomic sync, its a simple way to keep things organized (but very simple compared to mutex)
	// Research atomic functions in C++ to gain a better understanding.
}

// Private methods
func atomicCounters() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}

func regularCounters() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops++
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}
