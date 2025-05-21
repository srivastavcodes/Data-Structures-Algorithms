package main

import "math"

func maxPathSum(root *TreeNode) int {
	maxsum := math.MinInt
	helper124(root, &maxsum)
	return maxsum
}

func helper124(root *TreeNode, maxsum *int) int {
	if root == nil {
		return 0
	}
	left := helper124(root.Left, maxsum)
	right := helper124(root.Right, maxsum)

	left = max(left, 0)
	right = max(right, 0)

	pathSum := left + right + root.Val
	*maxsum = max(*maxsum, pathSum)
	return max(left, right) + root.Val
}
