package main

import "fmt"

// Go supports points
// If you want a good explination, I would suggest this video: https://www.youtube.com/watch?v=DTxHyVn0ODg&t=645s
// * dereferences the pointer, as well as signifies it is a pointer
// & gets the memory address of the value
func zeroval(ival int) {
	ival = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i) // this is a pass by reference, so the i value that was declared changes
	fmt.Println(i)

	fmt.Println(&i)
}
