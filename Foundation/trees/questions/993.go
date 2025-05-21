package main

func isCousins(root *TreeNode, x int, y int) bool {
	xx := findNode(root, x)
	yy := findNode(root, y)

	return level(root, xx, 0) == level(root, yy, 0) && !areSiblings(root, xx, yy)
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	node := findNode(root.Left, val)
	if node != nil {
		return node
	}
	return findNode(root.Right, val)
}

func level(root *TreeNode, val *TreeNode, lvl int) int {
	if root == nil {
		return 0
	}
	if root == val {
		return lvl
	}
	l := level(root.Left, val, lvl+1)
	if l != 0 {
		return l
	}
	return level(root.Right, val, lvl+1)
}

func areSiblings(root *TreeNode, x *TreeNode, y *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Left == x && root.Right == y ||
		root.Right == x && root.Left == y ||
		areSiblings(root.Left, x, y) || areSiblings(root.Right, x, y)
}
