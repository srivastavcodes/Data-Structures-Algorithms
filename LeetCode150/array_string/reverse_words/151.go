package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "the sky is blue"
	output := reverseWords(input)
	fmt.Println(output)
}

func reverseWords(input string) string {
	words := strings.Fields(input)

	low, high := 0, len(words)-1

	for low < high {
		words[low], words[high] = words[high], words[low]
		low++
		high--
	}
	return strings.TrimSpace(strings.Join(words, " "))
}
