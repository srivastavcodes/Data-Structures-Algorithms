package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	minIdx, maxIdx, culIdx := -1, -1, -1

	var result int64
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			culIdx = i
		}
		if nums[i] == minK {
			minIdx = i
		}
		if nums[i] == maxK {
			maxIdx = i
		}
		smaller := min(minIdx, maxIdx)

		temp := smaller - culIdx
		if temp < 0 {
			temp = 0
		}
		result += int64(temp)
	}
	return result
}
