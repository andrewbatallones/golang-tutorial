package sections

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Things to note
// if you want to pass the waitgroup, pass by reference
func WaitGrouops() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Because this is a reference to the for loop, if you pass this down to a concurrent goroutine,
		// you'll get the last value passed.
		i := i

		// can learn more on this from Closures
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// This basically blocks the code from continueing until all the waitgroups becomes 0
	wg.Wait()

	fmt.Println("All workers done.")
}
