package main

import "fmt"

func main() {
	strs := "abciiidef"
	k := 3

	count := maxVowels(strs, k)
	fmt.Println(count)

	input := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	output := longestOnes(input, 3)
	fmt.Println(output)
}
