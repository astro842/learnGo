package main

import "fmt"

func updateString(s []int) {
	s[0] = 100
}

func main() {
    //slice 是数组的映射(view)
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:4]
	s2 := arr[:]

	fmt.Println("s1:",s1)
	fmt.Println("s2:",s2)
	fmt.Println("-------")
	updateString(s1)
	//updateString(s2)
	fmt.Println("s1:",s1)
	fmt.Println("arr:",arr)
	fmt.Println("----append")

	s3 := arr[2:7]
	fmt.Println("s3:",s3)
	s4 := append(s3,10)
	s5 := append(s4,100)
	fmt.Println("s4:",s4)
	fmt.Println("s5:",s5)
	fmt.Println("arr:",arr)
}
