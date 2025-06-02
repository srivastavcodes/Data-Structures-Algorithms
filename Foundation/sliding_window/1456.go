package main

func maxVowels(str string, k int) int {
	left, right, result := 0, 0, 0

	count := 0
	for right < len(str) {
		if isVowels(str[right]) {
			count++
		}
		if right-left+1 == k {
			result = max(result, count)
			if isVowels(str[left]) {
				count--
			}
			left++
		}
		right++
	}
	return result
}

func isVowels(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
