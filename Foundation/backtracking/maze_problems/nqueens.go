package main

import "fmt"

func solveNQueens(queens int) [][]string {
	board := make([][]bool, queens)
	for i, _ := range board {
		board[i] = make([]bool, queens)
	}
	return buildAndSolveNQueens(board, 0)
}

func buildAndSolveNQueens(board [][]bool, rows int) [][]string {
	if rows == len(board) {
		result := convertBoardToStrings(board)
		return [][]string{result}
	}
	strs := make([][]string, 0)
	for col := 0; col < len(board); col++ {
		if isSafe(board, rows, col) {
			board[rows][col] = true
			strs = append(strs, buildAndSolveNQueens(board, rows+1)...)
			board[rows][col] = false
		}
	}
	return strs
}

func isSafe(board [][]bool, rows int, col int) bool {
	// check if queen present above
	for i := 0; i < rows; i++ {
		if board[i][col] {
			return false
		}
	}
	// check if present diagonally left
	leftMax := min(rows, col)
	for i := 1; i <= leftMax; i++ {
		if board[rows-i][col-i] {
			return false
		}
	}
	// check if present diagonally right
	rightMax := min(rows, len(board)-col-1)
	for i := 1; i <= rightMax; i++ {
		if board[rows-i][col+i] {
			return false
		}
	}
	return true
}

func display(board [][]bool) {
	for _, row := range board {
		for _, element := range row {
			if element {
				fmt.Printf("[Q]")
			} else {
				fmt.Printf("[ ]")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func convertBoardToStrings(board [][]bool) []string {
	result := make([]string, len(board))
	for i := 0; i < len(board); i++ {
		row := ""
		for j := range len(board[i]) {
			if board[i][j] {
				row += "Q"
			} else {
				row += "."
			}
		}
		result[i] = row
	}
	return result
}
