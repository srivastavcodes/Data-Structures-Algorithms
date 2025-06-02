package main

import "fmt"

func main() {
	strs := "abciiidef"
	k := 3
	count := maxVowels(strs, k)
	fmt.Println(count)

	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	output := maxSlidingWindow(input, k)
	fmt.Println(output)
}
