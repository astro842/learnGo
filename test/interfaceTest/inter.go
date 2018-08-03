package main

import "fmt"

type Animal interface {
	Speak() string
}

type Humen interface {
	Speak() string
	Sing() string
}

type Dog struct {
}

func (d Dog) Speak() string {

	return "Dog speak"
}


type Cat struct {
}

func (c Cat) Speak() string {

	return "Cat speak"
}

func (c Cat) Sing() string {

	return "Cat sing"
}

func main() {

	var a Animal = Dog{}
	fmt.Println(a.Speak())

	var c Humen = Cat{}
	fmt.Println(c.Speak())
	fmt.Println(c.Sing())
	}
