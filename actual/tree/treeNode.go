package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}
//结构的方法（类似于 Java类对象里的方法）
// (tree treeNode) 相当于this
//结构的方法与普通方法一样 只是调用形式不一样 @see: printTree1()
func (tree Node) PrintTree(){
   fmt.Println(tree.Value)
}


func PrintTree1(tree Node) {
	fmt.Print(tree.Value)
}


