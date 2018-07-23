package main

import "fmt"

func updateString(s []int) {
	s[0] = 100;
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("-------")
	updateString(s1)
	//updateString(s2)

	fmt.Println(s1)
	//fmt.Println(s2)
	fmt.Println(arr)

}
