package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4, 2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}
	curr := 0
	for {
		slow = nums[slow]
		curr = nums[curr]

		if curr == slow {
			return slow
		}
	}
}
