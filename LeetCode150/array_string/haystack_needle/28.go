package haystack_needle

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	for i := range len(haystack) + 1 - len(needle) {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
