package main

import "sort"

func subsetsWithDup(digits []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(digits)
	helper90(digits, &result, []int{}, 0)
	return result
}

func helper90(original []int, result *[][]int, subset []int, idx int) {
	if idx >= len(original) {
		subCopy := make([]int, len(subset))
		copy(subCopy, subset)
		*result = append(*result, subCopy)
		return
	}
	helper90(original, result, append(subset, original[idx]), idx+1)
	for idx+1 < len(original) && original[idx] == original[idx+1] {
		idx++
	}
	helper90(original, result, subset, idx+1)
}
