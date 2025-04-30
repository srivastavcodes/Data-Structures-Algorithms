package main

func main() {}

func dailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	var stack []int

	for i, curr := range temps {
		for len(stack) > 0 && curr > temps[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, i)
	}
	return result
}
