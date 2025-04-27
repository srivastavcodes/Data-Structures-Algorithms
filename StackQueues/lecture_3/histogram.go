package main

func largestRectangleArea(heights []int) int {
	var stack []int
	var maxArea int

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {

			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nse, pse := i, 0

			if len(stack) > 0 {
				pse = -1
			} else {
				pse = stack[len(stack)-1]
			}
			maxArea = max(maxArea, heights[curr]*(nse-pse-1))
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		nse, pse := len(heights), 0

		curr := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		if len(stack) > 0 {
			pse = -1
		} else {
			pse = stack[len(stack)-1]
		}
		maxArea = max(maxArea, heights[curr]*(nse-pse-1))
	}
	return maxArea
}
