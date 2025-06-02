package main

import "fmt"

func main() {
	strs := "pwwkew"
	res := lengthOfLongestSubstrings(strs)
	fmt.Println(res)
}

func lengthOfLongestSubstring(strs string) int {
	set := make(map[byte]bool)
	l, res := 0, 0

	for i := 0; i < len(strs); i++ {
		for set[strs[i]] {
			delete(set, strs[l])
			l++
		}
		set[strs[i]] = true
		if i-l+1 > res {
			res = i - l + 1
		}
	}
	return res
}

func lengthOfLongestSubstrings(strs string) int {
	record := make(map[rune]int)
	l, res := 0, 0

	for i, char := range strs {
		if idx, ok := record[char]; ok {
			l = max(l, idx+1)
		}
		record[char] = i
		res = max(res, i-l+1)
	}
	return res
}
