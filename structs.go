package main

import "fmt"

// stuct is a typed collection
// structs are mutable
type person struct {
	name string
	age  int
}

// You can attach methods to structs
type rect struct {
	width, height int
}

// You can pass a method by value or by reference
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
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
	fmt.Println("---")

	r := rect{42, 42}
	fmt.Println(r.area())
	fmt.Println(r.perim())
}
