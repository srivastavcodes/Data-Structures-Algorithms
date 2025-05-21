package main

func pathExists(root *TreeNode, arr []int) bool {
	if root == nil {
		return len(arr) == 0
	}
	return helperPathExists(root, arr, 0)
}

func helperPathExists(root *TreeNode, arr []int, idx int) bool {
	if root == nil {
		return false
	}
	if idx > len(arr)-1 || root.Val != arr[idx] {
		return false
	}
	if root.Left == nil && root.Right == nil && idx == len(arr)-1 {
		return true
	}
	return helperPathExists(root.Left, arr, idx+1) || helperPathExists(root.Right, arr, idx+1)
}
