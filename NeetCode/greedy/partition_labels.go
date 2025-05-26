package main

func partitionLabels(strs string) []int {
	lastIndex := make(map[rune]int)
	for i, s := range strs {
		lastIndex[s] = i
	}
	result := make([]int, 0)
	size, end := 0, 0
	for i, s := range strs {
		size++
		if lastIndex[s] > end {
			end = lastIndex[s]
		}
		if i == end {
			result = append(result, size)
			size = 0
		}
	}
	return result
}
