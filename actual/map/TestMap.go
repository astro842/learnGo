package main

import "fmt"

func main() {

	var mm = make(map[string]string)
	fmt.Println(mm)
	m := map[string]string{
		"name":  "aaa",
		"name1": "aaa1",
		"name2": "aaa2",
		"name3": "aaa3",
	}
	fmt.Println(m)
	fmt.Println("---遍历")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("-----get value")

	if s, c := m["name11"]; c {
		fmt.Println(s)
	} else {
		fmt.Println("key does not exist")
	}
	fmt.Println("-----delete value")
	if name, c := m["name"]; c {
		fmt.Println(name, c)
		delete(m, "name")
		fmt.Println(m)
	}
}
