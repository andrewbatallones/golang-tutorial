package main

import "fmt"

// stuct is a typed collection
// structs are mutable
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Amanda", age: 19})

	fmt.Println(*newPerson("Testington"))
}
