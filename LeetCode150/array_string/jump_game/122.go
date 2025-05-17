package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 0}
	output := canJump(nums)
	fmt.Println(output)
}

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
