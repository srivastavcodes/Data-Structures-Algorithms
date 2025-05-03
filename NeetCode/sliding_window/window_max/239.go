package main

func main() {}

func maxSlidingWindow(nums []int, k int) []int {
	var output, deque []int
	l, r := 0, 0

	for r < len(nums) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[r] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, r)
		if l > deque[0] {
			deque = deque[1:]
		}
		if (r + 1) >= k {
			output = append(output, nums[deque[0]])
			l++
		}
		r++
	}
	return output
}
