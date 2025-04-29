package main

import "unicode"

func main() {}

func isPalindrome(str string) bool {
	left, right := 0, len(str)-1

	for left < right {
		for left < right && !isAlphaNum(str[left]) {
			left++
		}
		for left < right && !isAlphaNum(str[right]) {
			right--
		}
		if unicode.ToLower(rune(str[left])) != unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c))
}
