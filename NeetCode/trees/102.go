package main

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		values := make([]int, 0)
		for range queue {
			currNode := queue[0]
			queue = queue[1:]

			values = append(values, currNode.Val)
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}
		result = append(result, values)
	}
	return result
}
