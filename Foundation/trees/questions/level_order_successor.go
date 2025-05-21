package main

func LevelOrderSuccessor(key int, root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lnode := queue[0]
		queue = queue[1:]

		if lnode.Left != nil {
			queue = append(queue, lnode.Left)
		}
		if lnode.Right != nil {
			queue = append(queue, lnode.Right)
		}
		if lnode.Val == key {
			break
		}
	}
	return queue[0]
}
