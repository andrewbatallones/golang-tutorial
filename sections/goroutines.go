package sections

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// Basically goroutines runs code on a different thread
func GoRoutines() {
	f("direct")

	go f("goroutine")

	// Setup a goroutine in an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
