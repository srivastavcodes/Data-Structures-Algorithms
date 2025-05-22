package main

import "fmt"

func main() {
	fmt.Printf("Maze problems\n")

	totalPaths := countPaths(3, 3)
	fmt.Printf("No. of ways:: %d\n", totalPaths)

	pathValues := returnPaths("", 3, 3)
	fmt.Printf("Paths to last cube:: %s\n", pathValues)

	withDiagonal := pathsWithDiagonal("", 3, 3)
	fmt.Printf("Paths with diagonal:: %s\n", withDiagonal)

	board := [][]bool{
		{true, true, true},
		{true, false, true},
		{true, true, true},
	}
	withRestrictions := pathsWithRestriction("", board, 0, 0)
	fmt.Printf("Paths with restrictions:: %s\n", withRestrictions)

	allPaths := allPathsWithRestriction("", board, 0, 0)
	fmt.Printf("All Paths with restrictions:: %s\n", allPaths)
}
