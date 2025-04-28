package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c]
			sqIdx := (r/3)*3 + c/3

			if rows[r][val] || cols[c][val] || squares[sqIdx][val] {
				return false
			}
			rows[r][val] = true
			cols[c][val] = true
			squares[sqIdx][val] = true
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	var column [9][9]int
	var row [9][9]int
	var box [3][3][9]int

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				k := board[i][j] - '1'
				column[j][k]++
				row[i][k]++
				box[i/3][j/3][k]++
				if column[j][k] > 1 || row[i][k] > 1 || box[i/3][j/3][k] > 1 {
					return false
				}
			}
		}
	}
	return true
}
