// Slices
// They are basically arrays, but dynamic
package main

import "fmt"

func main() {
	// 2 ways to create a slice
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s2 := []string{"a", "b", "c"}
	fmt.Println(s2)

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[0])
	fmt.Println("original len: ", len(s))

	// You can append the slice to a bigger size
	s = append(s, "d")
	s = append(s, "e", "f") // can pass multiple arguments

	// Can copy slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// can also take a slice from a slice
	// [1:2], [:5], [1:] are all acceptable
	fmt.Println("slice: ", s[1:2]) // low is include, high is excluded

	// Like arrays, slices can be multi-dimensional
	// They can be different lengths
	twoD := make([][]int, 2)
	twoD[0] = make([]int, 3)
	twoD[1] = make([]int, 2)
	fmt.Println(twoD)
}
