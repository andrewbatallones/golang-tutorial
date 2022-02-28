// This should cover up till:
//   Switch statements

package main

import (
	"fmt"
	"math"
	"time"
)

// You can declar constants with the 'cost' keyword
// You can also lock a var into a certain datatype
const global string = "global constant"

// Entrance for any app
func main() {
	const local float64 = 5000

	var text = "hello world"

	fmt.Println(text)
	fmt.Println(math.Sin(local))

	// For Loops
	i := 1 // shorthand for decalare/initialize variable
	// For is their only loop that they use
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// If statements (basic)
	// This is more of a general loop statement in other languages
	for j := 1; j <= 3; j++ {
		// This should print only even numbers
		// The brackets are required, but parentheses are not
		// There is no turnary operator in go :(
		if j%2 == 0 {
			fmt.Println(j)
		}
	}

	// Switch Statements
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Its a weekday...")
	}
}
