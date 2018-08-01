package main

import (
	"fmt"
	"learnGo/actual/tree"
)

//实现或者扩展别人写好的struct和方法
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.PrintTree()
}

func main() {

	var root tree.Node
	root = tree.Node{Value: 5}

	root.Left = &tree.Node{Value: 1}
	fmt.Println(root.Left.Value)

	root.Left.SetValue(100)
	fmt.Println(root.Left.Value)

	fmt.Println("-----------------")
	//root.printTree()
	//printTree1(root)

	var n *myTreeNode = &myTreeNode{&root}
	n.postOrder()

}
