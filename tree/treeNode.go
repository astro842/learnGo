package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}
//结构的方法（类似于 Java类对象里的方法）
// (tree treeNode) 相当于this
//结构的方法与普通方法一样 只是调用形式不一样 @see: printTree1()
func (tree TreeNode) printTree(){
   fmt.Println(tree.Value)
}
func (tree TreeNode) setValue(value int){
	tree.Value = value
}

func printTree1(tree TreeNode) {
	fmt.Println(tree.Value)
}

