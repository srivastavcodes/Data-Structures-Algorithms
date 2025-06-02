package main

// mine
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make(map[rune]int)
	for _, str := range s {
		record[str]++
	}
	for _, str := range t {
		if record[str] > 0 {
			record[str]--
		} else {
			return false
		}
	}
	return true
}

// optimized solution
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make([]int, 26)
	for i, _ := range s {
		record[s[i]-'a']++
		record[t[i]-'a']--
	}
	for _, r := range record {
		if r > 0 {
			return false
		}
	}
	return true
}
