package main

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	return helper230(root, k, &count).Val
}

func helper230(root *TreeNode, k int, count *int) *TreeNode {
	if root == nil {
		return root
	}
	left := helper230(root.Left, k, count)
	if left != nil {
		return left
	}
	*count++
	if *count == k {
		return root
	}
	return helper230(root.Right, k, count)
}
