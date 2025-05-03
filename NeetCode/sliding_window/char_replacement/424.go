package main

func main() {}

func characterReplacement(strs string, k int) int {
	result := 0
	charset := make(map[byte]struct{})

	for i := 0; i < len(strs); i++ {
		charset[strs[i]] = struct{}{}
	}
	for char := range charset {
		count, l := 0, 0
		for r := 0; r < len(strs); r++ {
			if strs[r] == char {
				count++
			}
			for (r-l+1)-count > k {
				if strs[l] == char {
					count--
				}
				l++
			}
			result = max(result, r-l+1)
		}
	}
	return result
}
