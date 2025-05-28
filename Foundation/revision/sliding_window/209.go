package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum, left, right := 0, 0, 0
	subarr := math.MaxInt

	for right < len(nums) {
		sum += nums[right]

		for sum >= target {
			subarr = min(subarr, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if subarr == math.MaxInt {
		return 0
	}
	return subarr
}
