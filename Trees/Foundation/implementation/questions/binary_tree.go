package questions

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func (tn *TreeNode) Insert(value int) {
	if value == tn.Val {
		return
	}
	if value < tn.Val {
		if tn.Left == nil {
			tn.Left = &TreeNode{Val: value}
		} else {
			tn.Left.Insert(value)
		}
	} else {
		if tn.Right == nil {
			tn.Right = &TreeNode{Val: value}
		} else {
			tn.Right.Insert(value)
		}
	}
}

func Traversal(node *TreeNode, indent string) {
	if node == nil {
		return
	}
	Traversal(node.Right, indent+"\t")
	fmt.Print(indent, node.Val, "\n")
	Traversal(node.Left, indent+"\t")
}

func PreOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	PreOrderTraversal(node.Left)
	PreOrderTraversal(node.Right)
}

func InOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	InOrderTraversal(node.Left)
	fmt.Print(node.Val, " ")
	InOrderTraversal(node.Right)
}

func PostOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrderTraversal(node.Left)
	PostOrderTraversal(node.Right)
	fmt.Print(node.Val, " ")
}

func (tn *TreeNode) PrettyDisplay() {
	Traversal(tn, "")
	fmt.Println()
	PreOrderTraversal(tn)
	fmt.Println()
	InOrderTraversal(tn)
	fmt.Println()
	PostOrderTraversal(tn)
}
