package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {

			node := queue[0]
			queue = queue[1:]

			if !leftToRight {
				levelVals[levelSize-1-i] = node.Val
			} else {
				levelVals[i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		leftToRight = !leftToRight
		result = append(result, levelVals)
	}
	return result
}
