package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func Insert(node *Node, value int) *Node {
	if node == nil {
		return NewNode(value)
	}

	if value < node.Value {
		node.Left = Insert(node.Left, value)
	} else {
		node.Right = Insert(node.Right, value)
	}

	return node
}

func Search(node *Node, value int) *Node {
	if node == nil || node.Value == value {
		return node
	}

	if value < node.Value {
		return Search(node.Left, value)
	} else {
		return Search(node.Right, value)
	}
}

func FindMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

func FindMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func main() {
	root := NewNode(4)
	Insert(root, 2)
	Insert(root, 6)
	Insert(root, 1)
	Insert(root, 3)
	Insert(root, 7)

	fmt.Println(Search(root, 5))
	fmt.Println(Search(root, 3))
	fmt.Println(FindMax(root))
	fmt.Println(FindMin(root))

}
