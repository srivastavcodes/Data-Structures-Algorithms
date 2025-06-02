package main

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}
	if len(nums) < 2*k+1 {
		return result
	}
	window := 0
	left, index, right := 0, k, 2*k

	for i := left; i <= right; i++ {
		window += nums[i]
	}
	elements := 2*k + 1
	result[index] = window / elements

	index++
	right++

	for right < len(nums) {
		outOfWindow := nums[left]
		intoWindow := nums[right]

		window = window + intoWindow - outOfWindow
		result[index] = window / elements

		left++
		index++
		right++
	}
	return result
}
