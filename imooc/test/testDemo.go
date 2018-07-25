package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s)
    c[0] = 'a'
    s2 := string(c)
	fmt.Println(s2)
	fmt.Println(s)
}
