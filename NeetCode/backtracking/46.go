package main

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	helper46(nums, &result, []int{}, 0)
	return result
}

func helper46(original []int, result *[][]int, subs []int, idx int) {
	if idx >= len(original) {
		subCopy := make([]int, len(subs))
		copy(subCopy, subs)
		*result = append(*result, subCopy)
		return
	}
	for i := idx; i < len(original); i++ {
		original[i], original[idx] = original[idx], original[i]
		helper46(original, result, append(subs, original[idx]), idx+1)
		original[idx], original[i] = original[i], original[idx]
	}
}
