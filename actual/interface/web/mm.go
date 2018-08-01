package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

//--------------------------

func (h Human) Say() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing() {
	fmt.Println("human sing")
}

func (h Human) Eat() {
	fmt.Println("human eat")
}

//-----------------------
func (s Student) GoSchool() {
	fmt.Println("go to school")
}
func (e Employee) GoWork() {
	fmt.Println("go to work")
}
func (e Employee) Say(s string) {
	fmt.Println("emp----", s)
}

//定义interface
//-----------------------
type Boy interface {
	Say()
	Sing()
	Eat()
	GoSchool()
	Post(s string)
}
type Men interface {
	Say(s string)
	Sing()
	Eat()
	GoWork()
}

func main() {
	tom := Student{Human{"Tom", 18, "123"}, "center school"}
	rose := Student{Human{"Rose", 18, "1234"}, "center school"}
	john := Employee{Human{"Jhon", 28, "456"}, "center worker"}

	var i Boy
	i = tom

	fmt.Println("This is ", tom.name)
	i.Say()
	//i.GoSchool()

	fmt.Println("---------------------")
	fmt.Println("This is ", john.name)
	john.Say("john - saying")
	john.GoWork()
	fmt.Println("---------------------")
	x := make([]Boy, 2)
	x[0], x[1] = tom, rose
	for _, value := range x{
		value.Say()
	}


}
