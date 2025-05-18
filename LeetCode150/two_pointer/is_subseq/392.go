package main

import "fmt"

func main() {
	s := "axc"
	t := "ahbgdc"
	fmt.Printf("output: %t", isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	l, r := 0, 0
	for range t {
		if s[l] == t[r] {
			l++
		}
		r++
		if l >= len(s) {
			return true
		}
	}
	return false
}
