package sections

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// One struct embeds another struct
// Doesn't require a datatype
type container struct {
	base
	str string
}

func Embedding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "Some string",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num: ", co.base.num)

	// the container struct inherits base's values and methods
	fmt.Println("describe: ", co.describe())

	// Can also embed an interface into a struct
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer: ", d.describe())
}
