package main

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool)
	left, right := 0, 0

	for right < len(nums) {
		if abs(left-right) > k {
			delete(set, nums[left])
			left++
		}
		if set[nums[right]] {
			return true
		}
		set[nums[right]] = true
		right++
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
