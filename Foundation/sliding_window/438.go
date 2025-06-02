package main

func findAnagrams(strs string, pattern string) []int {
	record := make([]int, 26)
	result := make([]int, 0)

	for i := range pattern {
		record[pattern[i]-'a']++
	}
	left, right := 0, 0

	for right < len(strs) {
		record[strs[right]-'a']--
		if right-left+1 == len(pattern) {
			if allZero(record) {
				result = append(result, left)
			}
			record[strs[left]-'a']++
			left++
		}
		right++
	}
	return result
}

func allZero2(record []int) bool {
	for i := range record {
		if record[i] != 0 {
			return false
		}
	}
	return true
}
