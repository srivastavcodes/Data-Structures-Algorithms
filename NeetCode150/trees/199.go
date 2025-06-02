package main

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i, curr := range queue {
			queue = queue[1:]
			if i == size-1 {
				result = append(result, curr.Val)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return result
}
