package main

func countOccurrence(pattern string, strs string) int {
	record := make([]int, 26)

	for i := range pattern {
		record[pattern[i]-'a']++
	}
	count, left, right := 0, 0, 0

	for right < len(strs) {
		record[strs[right]-'a']--
		if right-left+1 == len(pattern) {
			if allZero(record) {
				count++
			}
			record[strs[left]-'a']++
			left++
		}
		right++
	}
	return count
}

func allZero(record []int) bool {
	for i := range record {
		if record[i] != 0 {
			return false
		}
	}
	return true
}
