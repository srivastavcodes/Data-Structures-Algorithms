package main

func AverageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	queue := []*TreeNode{root}

	if root == nil {
		return result
	}
	for len(queue) > 0 {
		levelSize := len(queue)
		var average float64
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			average += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		average = average / float64(levelSize)
		result = append(result, average)
	}
	return result
}
