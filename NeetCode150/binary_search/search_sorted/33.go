package main

func main() {}

func search(nums []int, tgt int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if tgt == nums[mid] {
			return mid
		}
		if nums[l] <= nums[mid] {
			if tgt > nums[mid] || tgt < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if tgt < nums[mid] || tgt > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
