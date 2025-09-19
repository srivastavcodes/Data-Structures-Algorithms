package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "0P"
	out := isPalindrome(s)
	fmt.Println(out)
}

func isPalindrome(strs string) bool {
	left, right := 0, len(strs)-1

	for left < right {
		for left < right && !alphaNumeric(rune(strs[left])) {
			left++
		}
		for left < right && !alphaNumeric(rune(strs[right])) {
			right--
		}
		if unicode.ToLower(rune(strs[left])) != unicode.ToLower(rune(strs[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func alphaNumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}
