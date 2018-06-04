package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const path  = "e://copy.txt"
	contents,err :=ioutil.ReadFile(path)
	if err!=nil {
		fmt.Print(err)
	}else {
		fmt.Printf("%s\n ",contents)
	}
}
