package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

// todo -> sorted array to binary tree, do it!

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

func PreOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	var result []int
	preorderTraverse(node, &result)
	return result
}

func preorderTraverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preorderTraverse(node.Left, result)
	preorderTraverse(node.Right, result)
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
	values := []interface{}{3, 9, 11, 24, 18, 19, 4}
	root := BuildTree(values)

	root.PrettyDisplay()

	result := levelOrder(root)
	fmt.Printf("level order:: \t%d\n", result)

	zigzag := zigzagLevelOrder(root)
	fmt.Printf("zigzag:: \t%d\n", zigzag)

	bottom := levelOrderBottom(root)
	fmt.Printf("bottom up:: \t%d\n", bottom)

	rightView := rightSideView(root)
	fmt.Printf("right view:: \t%d\n", rightView)

	symmetric := isSymmetric(root)
	fmt.Printf("symmetric:\t%t\n", symmetric)

	diameter := diameterOfBinaryTree(root)
	fmt.Printf("diameter:\t%d\n", diameter)

	inverted := invertTree(root)
	inverted.PrettyDisplay()

	preorder := PreOrderTraversal(root)
	fmt.Printf("preorder:\t%d\n", preorder)
}
