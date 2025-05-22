package main

func subsets(digits []int) [][]int {
	result := make([][]int, 0)
	helper78(&result, []int{}, digits, 0)
	return result
}

func helper78(result *[][]int, subset, digits []int, i int) {
	if i >= len(digits) {
		subCopy := make([]int, len(subset))
		copy(subCopy, subset)
		*result = append(*result, subCopy)
		return
	}
	helper78(result, subset, digits, i+1)
	helper78(result, append(subset, digits[i]), digits, i+1)
}
