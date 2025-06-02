package main

func main() {}

func searchMatrix(matrix [][]int, tgt int) bool {
	rows, cols := len(matrix), len(matrix[0])
	top, bot := 0, rows-1

	for top <= bot {
		row := (top + bot) / 2
		if tgt > matrix[row][cols-1] {
			top = row + 1
		} else if tgt < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}
	if !(top <= bot) {
		return false
	}
	row := (top + bot) / 2
	l, r := 0, cols-1
	for l <= r {
		mid := (l + r) / 2
		if tgt > matrix[row][mid] {
			l = mid + 1
		} else if tgt < matrix[row][mid] {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
