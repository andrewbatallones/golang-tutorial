package main

import "fmt"

// Go supports anonymous functions
// func() int is the return type
// This leads go into supporting closures
// A closure is just a function wrapped in another function (kind of like a namespace)
// can read more here: https://en.wikipedia.org/wiki/Closure_(computer_programming)
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
