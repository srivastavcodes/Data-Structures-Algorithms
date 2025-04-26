package main

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	stack := make([]int, 0, len(nums))
	for i := 0; i < 2*len(nums); i++ {
		i := i % len(nums)
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			ans[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
