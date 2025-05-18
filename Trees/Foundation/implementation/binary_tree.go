package main

import "fmt"

var root *TreeNode

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

func (tn *TreeNode) insert(value int) {
	if value < tn.Value {
		if tn.Left == nil {
			tn.Left = &TreeNode{Value: value}
		} else {
			tn.Left.insert(value)
		}
	}
	if value > tn.Value {
		if tn.Right == nil {
			tn.Right = &TreeNode{Value: value}
		} else {
			tn.Right.insert(value)
		}
	}
}

func InorderTraversal(node *TreeNode, indent string) {
	if node == nil {
		return
	}
	InorderTraversal(node.Right, indent+"\t")
	fmt.Print(indent, node.Value, "\n")
	InorderTraversal(node.Left, indent+"\t")
}

func prettyDisplay() {
	InorderTraversal(root, "")
}

func main() {
	root = &TreeNode{Value: 10}
	root.insert(5)
	root.insert(15)
	root.insert(3)
	root.insert(7)
	root.insert(12)
	root.insert(18)
	root.insert(9)
	root.insert(22)
	root.insert(32)
	root.insert(2)
	root.insert(11)

	prettyDisplay()
}
