package main

import "fmt"

func main() {
	strs := "cbaebabacd"
	pattern := "abc"

	count := findAnagrams(strs, pattern)
	fmt.Println(count)

	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	subarr := minSubArrayLen(11, nums)
	fmt.Println(subarr)
}
