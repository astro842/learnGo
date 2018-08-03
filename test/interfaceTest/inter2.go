package main

import (
	"fmt"
)

//接口定义方法
type name interface {
	area() int
	hello() string
}

//定义一个类型
type rect struct {
	height, weight int
}

//类型的方法与 接口的方法同名
//类型需要实现 接口的所有方法才属于那个方法
func (r rect) area() int {
	return r.height * r.weight
}
func (r rect) hello() string {
	return "hello"
}

func main() {
	var r name = rect{3, 4}
	fmt.Println(r.area())
	fmt.Println(r.hello())

}
