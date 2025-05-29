package main

import "fmt"

func main() {
	strs := "abciiidef"
	k := 3

	count := maxVowels(strs, k)
	fmt.Println(count)

	input := []int{1, 2, 3, 4, 5}
	output := getAverages(input, k)
	fmt.Println(output)
}
