package main

func isBalanced(root *TreeNode) bool {
	result := true
	helper110(root, &result)
	return result
}

func helper110(root *TreeNode, result *bool) int {
	if root == nil {
		return 0
	}
	left := helper110(root.Left, result)
	right := helper110(root.Right, result)
	if abs(left-right) > 1 {
		*result = false
	}
	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
