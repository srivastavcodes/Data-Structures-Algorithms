package main

func subsets(digits []int) [][]int {
	if len(digits) == 0 {
		return [][]int{{}}
	}
	return getSubsets([]int{}, digits)
}

func getSubsets(sub []int, original []int) [][]int {
	if len(original) == 0 {
		return [][]int{append([]int{}, sub...)}
	}
	digit := original[0]

	appended := append(append([]int{}, sub...), digit)
	left := getSubsets(sub, original[1:])
	right := getSubsets(appended, original[1:])

	return append(left, right...)
}
