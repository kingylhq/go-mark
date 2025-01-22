package example

import "fmt"

type rect2 struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.   面积
func (r *rect2) area() int {
	return r.width * r.height
}

// 周长
// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect2) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {

	r := rect{width: 10, height: 5}
	fmt.Printf("长方形 width = %v, height = %v \n", r.width, r.height)

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())

	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
