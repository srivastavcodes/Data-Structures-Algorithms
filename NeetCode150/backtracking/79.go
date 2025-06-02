package main

// todo -> draw it and understand the recursion

func exist(board [][]byte, word string) bool {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if helper79(board, word, r, c, 0) {
				return true
			}
		}
	}
	return false
}

func helper79(board [][]byte, word string, r, c, i int) bool {
	rows, cols := len(board), len(board[0])
	if i == len(word) {
		return true
	}
	if r < 0 || c < 0 || r >= rows || c >= cols ||
		board[r][c] != word[i] || board[r][c] == '#' {
		return false
	}
	curr := board[r][c]
	board[r][c] = '#'

	result := helper79(board, word, r+1, c, i+1) ||
		helper79(board, word, r-1, c, i+1) ||
		helper79(board, word, r, c+1, i+1) || helper79(board, word, r, c-1, i+1)

	board[r][c] = curr
	return result
}
