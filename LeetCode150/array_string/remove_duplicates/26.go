package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 2}
	res := removeDuplicates(nums)
	fmt.Printf("remove duplicates: %d", res)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	lv, l := math.MinInt, 0
	for i := range nums {
		if nums[i] != lv {
			nums[l] = nums[i]
			l++
		}
		lv = nums[i]
	}
	return l
}
