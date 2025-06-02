package main

import (
	"math"
)

func minWindow(strs string, sub string) string {
	length := len(strs)
	if len(sub) > length {
		return ""
	}
	record := make(map[byte]int)
	for i := range sub {
		record[sub[i]]++
	}
	requiredCount := len(sub)
	left, right := 0, 0

	windowSize, startIdx, lastIdx := math.MaxInt, 0, 0

	for right < len(strs) {
		char := strs[right]

		if record[char] > 0 {
			requiredCount--
		}
		record[char]--
		for requiredCount == 0 {
			currWindowSize := right - left + 1

			if currWindowSize < windowSize {
				windowSize = currWindowSize
				startIdx = left
				lastIdx = right
			}
			record[strs[left]]++
			if record[strs[left]] > 0 {
				requiredCount++
			}
			left++
		}
		right++
	}
	if windowSize == math.MaxInt {
		return ""
	}
	return strs[startIdx : lastIdx+1]
}
