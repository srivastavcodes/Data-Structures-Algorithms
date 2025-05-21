package main

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

func (tn *TreeNode) insert(value int) {
	if value == tn.Value {
		return
	}
	if value < tn.Value {
		if tn.Left == nil {
			tn.Left = &TreeNode{Value: value}
		} else {
			tn.Left.insert(value)
		}
	} else {
		if tn.Right == nil {
			tn.Right = &TreeNode{Value: value}
		} else {
			tn.Right.insert(value)
		}
	}
}

func Traversal(node *TreeNode, indent string) {
	if node == nil {
		return
	}
	Traversal(node.Right, indent+"\t")
	fmt.Print(indent, node.Value, "\n")
	Traversal(node.Left, indent+"\t")
}

func PreOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	PreOrderTraversal(node.Left)
	PreOrderTraversal(node.Right)
}

func InOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	InOrderTraversal(node.Left)
	fmt.Print(node.Value, " ")
	InOrderTraversal(node.Right)
}

func PostOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrderTraversal(node.Left)
	PostOrderTraversal(node.Right)
	fmt.Print(node.Value, " ")
}

func (tn *TreeNode) prettyDisplay() {
	Traversal(tn, "")
}

func main() {
	root := &TreeNode{Value: 5}
	root.insert(4)
	root.insert(11)
	root.insert(3)
	root.insert(6)
	root.insert(7)

	root.prettyDisplay()
}
