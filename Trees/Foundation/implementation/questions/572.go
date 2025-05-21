package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return helper572(root, subRoot)
}

func helper572(p, q *TreeNode) bool {
	if p == nil {
		return false
	}
	if sameTree(p, q) {
		return true
	}
	return helper572(p.Left, q) || helper572(p, q.Right)
}

func sameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}
