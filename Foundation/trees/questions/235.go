package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestorItr(root, p, q *TreeNode) *TreeNode {
	for {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
}
