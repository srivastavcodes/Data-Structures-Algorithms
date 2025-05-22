package main

func countPaths(rows int, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	left := countPaths(rows-1, cols)
	right := countPaths(rows, cols-1)
	return left + right
}
