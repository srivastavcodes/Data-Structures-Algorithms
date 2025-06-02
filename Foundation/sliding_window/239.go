package main

// todo -> LeetCode 1425

func maxSlidingWindow(nums []int, k int) []int {
	result, deque := make([]int, 0), make([]int, 0)

	for right := 0; right < len(nums); right++ {
		for len(deque) > 0 && deque[0] <= right-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[right] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)
		if right >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
