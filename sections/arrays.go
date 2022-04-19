package sections

import "fmt"

func Arrays() {

	// You need to initialize with a set len
	// This means you prolly can't add to array
	var a [5]int
	fmt.Println("emp: ", a)
	fmt.Println("len: ", len(a))

	// Assignment
	a[0] = 1
	fmt.Println(a)

	// You can do multi dimensional arrays
	var multiDemArray [2][3]int
	fmt.Println("Multi array: ", multiDemArray)
}
