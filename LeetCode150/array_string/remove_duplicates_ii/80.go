package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3}
	res := removeDuplicates(nums)
	fmt.Printf("remove duplicates: %d", res)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	lv, l, count := math.MinInt, 0, 0
	for i := range nums {
		if nums[i] != lv || count < 2 {
			if nums[i] != lv {
				count = 0
			}
			nums[l] = nums[i]
			l++
			count++
		}
		lv = nums[i]
	}
	fmt.Printf("final array: %v", nums)
	return l
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	var slow = 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
