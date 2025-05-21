package main

func sumNumbers(root *TreeNode) int {
	return helper129(root, 0)
}

func helper129(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return helper129(root.Left, sum) + helper129(root.Right, sum)
}
