package main

import "fmt"

func f(a int)  {
	a=a+1
}
func f3(p int) int {
	return p+1
}

func f4(p *int)  {
	*p = *p +1
}

func f2(p *int)  {
	fmt.Println("p =",p)
	fmt.Println("*p =",*p)
	fmt.Println("&p =",&p)
}

func main() {
	a := 5
	fmt.Println("a = ", a)
	f2(&a)
	fmt.Println("f2(),a =",a)

	//fmt.Println("f3()后，a = ",f3(a))
	//f4(&a)
	//fmt.Println("f4()后，a = ", a)

}
