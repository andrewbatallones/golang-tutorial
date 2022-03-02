// This should cover up till:
//   Switch statements

// Variables
// You need to use every variable or else go will blow up with an error
// Lower case variables are bound to the package vs upper case variables are exposed beyond the package
// varName will have a packatge scope
// VarName will have a global scope
// Acronyms: They should all be uppercase for readability
// var targetUrl ==> (targetURL)

// This is kind of arbritrary
// Basically specifies which package the code belongs to
// main is special becuase it serves as the entry point package
package main

import (
	"fmt"
	"math"
	"time"
)

// You can declar constants with the 'cost' keyword
// You can also lock a var into a certain datatype
// Cannot use := syntax here
const global string = "global constant"

// Shadowing: variable with the same name, but one at outer scope, other inner scope
var shadow string = "hi"

// Entrance for any app
func main() {
	const local float64 = 5000 // If you need more control

	var text = "hello world"

	// You can also declare variables this way too
	// One design choice would be if these variables are related in one way or another
	var (
		one = "one"
		// etc..
	)

	// This variable should print instead
	var shadow string = "I should print instead"

	// Typecasting
	var firstNum = 42
	var secondNum = float64(firstNum)
	fmt.Println(secondNum)
	// Gotcha: you'll need to bring in the "strconv" package to convert numbers to strings and vise versa

	fmt.Println(text)
	fmt.Println(one)
	fmt.Println(math.Sin(local))
	fmt.Println(shadow)

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
