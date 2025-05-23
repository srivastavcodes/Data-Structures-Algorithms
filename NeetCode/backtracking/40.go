package main

import "sort"

func combinationSum2(candidates []int, tgt int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	helper40(candidates, tgt, &result, []int{}, 0, 0)
	return result
}

func helper40(original []int, k int, result *[][]int, sub []int, sum, i int) {
	if sum == k {
		subCopy := make([]int, len(sub))
		copy(subCopy, sub)
		*result = append(*result, subCopy)
		return
	}
	if i >= len(original) || sum > k {
		return
	}
	helper40(original, k, result, append(sub, original[i]), sum+original[i], i+1)
	for i+1 < len(original) && original[i] == original[i+1] {
		i++
	}
	helper40(original, k, result, sub, sum, i+1)
}
