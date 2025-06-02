package main

func lengthOfLongestSubstring(strs string) int {
	charSet := make(map[byte]int)
	result, l := 0, 0

	for r := range strs {
		if idx, ok := charSet[strs[l]]; ok {
			l = max(l, idx+1)
		}
		charSet[strs[r]] = r
		result = max(result, r-l+1)
	}
	return result
}
