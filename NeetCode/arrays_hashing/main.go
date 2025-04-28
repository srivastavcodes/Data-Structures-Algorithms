package main

import "fmt"

func main() {
	_ = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ints := []int{4, 7, 2, 3, 9, 5, 6}

	output := longestConsecutive(ints)
	fmt.Println(output)
}
