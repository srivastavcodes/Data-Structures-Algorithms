package main

import "fmt"

func main() {
	nums := []int{-1}
	output := rotate(nums, 3)
	fmt.Println(output)
}

func rotate(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}
	k = k % len(nums)
	if k > 0 {
		rotateByIndex(nums, 0, len(nums)-k-1)
		rotateByIndex(nums, len(nums)-k, len(nums)-1)
		rotateByIndex(nums, 0, len(nums)-1)
	}
	return nums
}

func rotateByIndex(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
