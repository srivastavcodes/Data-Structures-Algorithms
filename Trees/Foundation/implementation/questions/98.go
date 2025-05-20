package main

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root *TreeNode, low *int, high *int) bool {
	if root == nil {
		return true
	}
	if low != nil && root.Val <= *low {
		return false
	}
	if high != nil && root.Val >= *high {
		return false
	}
	left := helper(root.Left, low, &root.Val)
	right := helper(root.Right, &root.Val, high)

	return left && right
}
