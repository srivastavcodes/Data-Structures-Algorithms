package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func BuildTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
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
}

func main() {
	values := []interface{}{3, 4, 8, 12, 14, 7, 11, nil, nil, 36, 9, nil, nil, 2, nil, nil, nil, nil, nil, nil, 44}
	root := BuildTree(values)

	root.PrettyDisplay()

	result := LevelOrder(root)
	average := AverageOfLevels(root)
	successor := LevelOrderSuccessor(14, root)
	fmt.Println(result)
	fmt.Println(average)
	fmt.Println(successor.Val)
}
