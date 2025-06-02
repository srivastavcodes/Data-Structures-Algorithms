package main

func combinationSum(candidates []int, tgt int) [][]int {
	result := make([][]int, 0)
	helper39(candidates, tgt, &result, []int{}, 0, 0)
	return result
}

func helper39(original []int, k int, result *[][]int, sub []int, sum, i int) {
	if sum == k {
		subCopy := make([]int, len(sub))
		copy(subCopy, sub)
		*result = append(*result, subCopy)
		return
	}
	if i >= len(original) || sum > k {
		return
	}
	helper39(original, k, result, sub, sum, i+1)
	helper39(original, k, result, append(sub, original[i]), sum+original[i], i)
}
