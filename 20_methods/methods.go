package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) x2() *rect {
	r.width *= 2
	r.height *= 2
	return r
}

func (r rect) x3() rect {
	r.width *= 3
	r.height *= 3
	return r
}

func (r rect) perim() int {
	return r.width + r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	fmt.Println("do x3: ", r.x3())
	fmt.Println("after x3: ", r)

	fmt.Println("do x2: ", r.x2())
	fmt.Println("after x2: ", r)
}
