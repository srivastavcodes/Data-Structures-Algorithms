package main

func main() {}

func maxArea(heights []int) int {
	left, right, result := 0, len(heights)-1, 0

	for left < right {
		area := min(heights[left], heights[right]) * (right - left)
		result = max(result, area)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
