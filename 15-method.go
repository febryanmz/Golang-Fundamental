package main

import "fmt"

type rect struct {
	width, height int
}

// method
func (r rect) area() int {
	return r.width * r.height
}

// method
func (b rect) perim() int {
	return 2*b.width + 2*b.height
}

func main() {
	kotak := rect{width: 10, height: 5}

	fmt.Println("area: ", kotak.area())
	fmt.Println("perim:", kotak.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

//febryan
