package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))

}
func test1() {
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}
func test2() {
	s1:= []int{2,3,4,5,6}
	printSlice(s1)
	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)
	fmt.Println("copy-----")
	copy(s2,s1)
	printSlice(s2)
	fmt.Println("delete------")

	s4 :=append(s2[:4],s2[5:]...)
	printSlice(s4)

}

func main() {
	//test1()
	test2()
}
