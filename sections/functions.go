package sections

import "fmt"

// Functions
// You must specify the data types of parameters
// if they are the same, you can omit the data types (except for the last one)
// You must have a return statement (does not return last expression)
// NOTE: go does not support function overloading and does not support user defined operators
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// Functions support multiple return values
func checkPairEven(a int, b int) (bool, bool) {
	return a%2 == 0, b%2 == 0
}

// Functions support varadic functions (limitless parameters)
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// Go supports recursions
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func Functions() {
	fmt.Println(plus(1, 2))
	fmt.Println(plusPlus(1, 2, 3))

	_, val := checkPairEven(3, 4)
	fmt.Println("second value: ", val)

	fmt.Println(sum(1, 2, 3, 4, 5))
}
