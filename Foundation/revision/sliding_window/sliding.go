package main

import "fmt"

func main() {
	strs := []int{1, 2, 3, 1, 2, 3}
	pattern := 2

	count := containsNearbyDuplicate(strs, pattern)
	fmt.Println(count)
}
