package main

// todo -> revise

func numSubarraysWithSum(nums []int, goal int) int {
	left, right, windowSum := 0, 0, 0

	result, zeroCount := 0, 0
	for right < len(nums) {
		windowSum += nums[right]

		for left < right && (nums[left] == 0 || windowSum > goal) {
			if nums[left] == 0 {
				zeroCount++
			} else {
				zeroCount = 0
			}
			windowSum -= nums[left]
			left++
		}
		if windowSum == goal {
			result += 1 + zeroCount
		}
		right++
	}
	return result
}
