package sections

import "fmt"

func Ranges() {
	nums := []int{1, 2, 3, 4}
	sum := 0

	// range will provide both an index and the iteration
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	// range on maps iterates both key/value pairs
	kvs := map[string]string{"a": "value1", "b": "value2"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Can also loop over just the keys
	for k := range kvs {
		fmt.Println("Key: ", k)
	}

	// Strings will loop over the unicode vs the actual character
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
