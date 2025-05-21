package main

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		l := queue[0]
		r := queue[1]
		queue = queue[2:]

		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left)
		queue = append(queue, r.Right)
		queue = append(queue, l.Right)
		queue = append(queue, r.Left)
	}
	return true
}
