package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}
//结构的方法（类似于 Java类对象里的方法）
// (tree treeNode) 相当于this
//结构的方法与普通方法一样 只是调用形式不一样 @see: printTree1()
func (tree treeNode) printTree(){
   fmt.Println(tree.value)
}
func (tree treeNode) setValue(value int){
	tree.value = value
}

func printTree1(tree treeNode) {
	fmt.Println(tree.value)
}

func main() {

	var root treeNode
	root = treeNode{value: 5}

	root.left = &treeNode{value:1}
	fmt.Println(root.left.value)

	root.left.setValue(100)
	fmt.Println(root.left.value)




	fmt.Println("-----------------")
	//root.printTree()
	//printTree1(root)

}
