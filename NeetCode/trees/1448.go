package main

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper1148(root, root.Val)
}

func helper1148(root *TreeNode, maxval int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val >= maxval {
		result = 1
	}
	maxval = max(maxval, root.Val)
	result += helper1148(root.Left, maxval)
	result += helper1148(root.Right, maxval)
	return result
}
