package sections

import (
	"fmt"

	"github.com/andrewbatallones/golang-tutorial/shapes"
)

// Interfaces are named collections of method signatures
// Almost think of them as interfaces in Java where you can pass in the
//    interface and as long as the struct follows the interface signatures, should produce good results
type geometry interface {
	Area() float64
	Perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func Interfaces() {
	r := shapes.Rect{3, 4}
	c := shapes.Circle{5}

	measure(r)
	measure(c)
}
