package main

import "fmt"

func main() {
	fmt.Printf("Maze problems\n")

	chess := make([][]bool, 4)
	for i := range chess {
		chess[i] = make([]bool, 4)
	}
	res := buildAndSolveNQueens(chess, 0)
	fmt.Println(res)
	fmt.Println()

	queensPlaced := solveNQueens(1)
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
