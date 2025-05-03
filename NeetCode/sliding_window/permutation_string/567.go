package main

import "fmt"

func main() {
	str1, str2 := "ab", "eidbaooo"
	output := checkInclusion(str1, str2)
	fmt.Println(output)
}

func checkInclusion(str1, str2 string) bool {
	if len(str1) > len(str2) {
		return false
	}
	s1Count := make([]int, 26)
	s2Count := make([]int, 26)
	for i := 0; i < len(str1); i++ {
		s1Count[str1[i]-'a']++
		s2Count[str2[i]-'a']++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	l := 0
	for r := len(str1); r < len(str2); r++ {
		if matches == 26 {
			return true
		}
		index := str2[r] - 'a'
		s2Count[index]++

		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]+1 == s2Count[index] {
			matches--
		}
		index = str2[l] - 'a'
		s2Count[index]--

		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]-1 == s2Count[index] {
			matches--
		}
		l++
	}
	return matches == 26
}
