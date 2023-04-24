package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// creating node
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// order is Root,Left,Right
func PreOrderTraversal(node *Node) {
	if node != nil {
		//for preorder
		fmt.Printf("%d ", node.Value)
		PreOrderTraversal(node.Left)
		PreOrderTraversal(node.Right)
	}
	//for inorder left,root,right
	//for postorder left,right,root
}

func main() {
	root := NewNode(5)
	root.Left = NewNode(3)
	root.Left.Left = NewNode(1)
	root.Right = NewNode(7)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(9)

	PreOrderTraversal(root)
}
