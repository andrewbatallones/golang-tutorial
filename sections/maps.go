package sections

import "fmt"

func Maps() {
	// initializing a map is similar to a slice
	// You use the make method, but you specify its a map and add the key type in the brackets
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2

	// Same with slices, you can initialize a map this way
	m2 := map[string]int{"k1": 1, "k2": 2}

	fmt.Println("map: ", m2)
	fmt.Println("map: ", m)
	fmt.Println("value1: ", m["k1"])

	// You ca remove a key by calling the delete method
	delete(m, "k2")
	fmt.Println("mod map: ", m)

	// Getting a value actually returns 2 values
	// the value itself, if the value was present
	// This is because Go uses 0 as the default value, so if a key doesn't exists, it returns 0
	// Another new term is the blank identifier _
	// just means we don't really need that value
	_, valuleExists := m["k2"]
	fmt.Println("value exists? ", valuleExists)
}
