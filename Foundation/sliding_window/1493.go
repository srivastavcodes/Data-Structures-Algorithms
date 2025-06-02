package main

// todo -> ***Similar Problem : LeetCode - 1004, 487, 485***

func longestSubarray(nums []int) int {
	zeroCount := 0
	maxLength := 0

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		maxLength = max(maxLength, right-left)
	}
	return maxLength
}

func longestSubarray2(nums []int) int {
	lastZeroIndex := 0
	longestSubarr := 0

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			left = lastZeroIndex + 1
			lastZeroIndex = right
		}
		longestSubarr = max(longestSubarr, right-left)
	}
	return longestSubarr
}
