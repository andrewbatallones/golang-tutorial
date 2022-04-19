package sections

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// Needs to be capitalize or else it will not be exported
type testStruct struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Json() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	// ... can go on for other data types

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "banana", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	dat, err := os.ReadFile("test.json")
	check(err)

	var testResponse testStruct
	uerr := json.Unmarshal(dat, &testResponse)

	if uerr != nil {
		fmt.Println(uerr)
	}

	fmt.Println("key1: ", testResponse.Key1)
	fmt.Println("key2: ", testResponse.Key2)
}
