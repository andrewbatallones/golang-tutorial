package sections

import (
	"fmt"
	"unicode/utf8"
)

// Strings are read-only slice of bytes
// Gonna have to read more about strings/runes here: https://go.dev/blog/strings
func Strings() {
	const s = "unchaning"
	const hello = "สวัสดี" // "Hello" in Thai

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println("---")

	// You can start seeing this when using a different language or special characters
	fmt.Println("Len: ", len(hello))
	fmt.Println("Rune count: ", utf8.RuneCountInString(hello))

	for i := 0; i < len(hello); i++ {
		fmt.Printf("%x ", hello[i])
	}
}
