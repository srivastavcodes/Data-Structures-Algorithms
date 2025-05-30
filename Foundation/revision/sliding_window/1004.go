package main

func longestOnes(nums []int, k int) int {
	result, zeroCount := 0, 0

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
