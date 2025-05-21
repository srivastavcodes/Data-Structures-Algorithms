package main

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	calcHeight(root, &diameter)
	return diameter - 1
}

func calcHeight(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := calcHeight(root.Left, diameter)
	right := calcHeight(root.Right, diameter)

	*diameter = max(*diameter, left+right+1)

	return max(left, right) + 1
}
