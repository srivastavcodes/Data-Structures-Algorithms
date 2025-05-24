package main

func isHeap(root *TreeNode) bool {
	index := 0
	count := countNodes(root)
	if isCBT(root, index, count) && isMaxHeap(root) {
		return true
	} else {
		return false
	}
}

func isMaxHeap(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Right == nil {
		return root.Value > root.Left.Value
	}
	return root.Value > root.Left.Value &&
		root.Value > root.Right.Value && isMaxHeap(root.Left) && isHeap(root.Right)
}

func isCBT(root *TreeNode, index int, count int) bool {
	if root == nil {
		return true
	}
	if index >= count {
		return false
	}
	return isCBT(root.Left, 2*index+1, count) &&
		isCBT(root.Right, 2*index+2, count)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
