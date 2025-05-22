package main

import "fmt"

func main() {
	fmt.Printf("Maze problems\n")

	queensPlaced := solveNQueens(5)
	fmt.Println(queensPlaced)
	fmt.Println()

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
