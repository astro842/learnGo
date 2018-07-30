package main

import "fmt"

type person struct {
	name string
	age int
}

func (r Rect) size() int {
	return r.width * r.height
}

type Rect struct {
	width  int
	height int
}
func main() {

	a := Rect{height: 1, width: 100}
	f := &Rect{111, 111}

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println("-------------------------")

	fmt.Println(f)
	fmt.Println(*f)
	fmt.Println(&f)
}
