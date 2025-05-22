package main

// without diagonal traversal
func returnPaths(pr string, rows, cols int) []string {
	if rows == 1 && cols == 1 {
		return []string{pr}
	}
	paths := make([]string, 0)
	if cols > 1 {
		paths = append(paths, returnPaths(pr+"\u2192", rows, cols-1)...)
	}
	if rows > 1 {
		paths = append(paths, returnPaths(pr+"\u2193", rows-1, cols)...)
	}
	return paths
}

// with diagonal traversal
func pathsWithDiagonal(pr string, rows, cols int) []string {
	if rows == 1 && cols == 1 {
		return []string{pr}
	}
	paths := make([]string, 0)
	if rows > 1 && cols > 1 {
		diagPaths := pathsWithDiagonal(pr+"\u2198", rows-1, cols-1)
		paths = append(paths, diagPaths...)
	}
	if cols > 1 {
		horizPaths := pathsWithDiagonal(pr+"\u2192", rows, cols-1)
		paths = append(paths, horizPaths...)
	}
	if rows > 1 {
		vertPaths := pathsWithDiagonal(pr+"\u2193", rows-1, cols)
		paths = append(paths, vertPaths...)
	}
	return paths
}

// maze with restrictions
func pathsWithRestriction(pr string, board [][]bool, rows, cols int) []string {
	paths := make([]string, 0)
	if rows == len(board)-1 && cols == len(board[0])-1 {
		return []string{pr}
	}
	if !board[rows][cols] {
		return paths
	}
	if rows < len(board)-1 && cols < len(board[0])-1 {
		diagPaths := pathsWithRestriction(pr+"\u2198", board, rows+1, cols+1)
		paths = append(paths, diagPaths...)
	}
	if cols < len(board[0])-1 {
		horizPaths := pathsWithRestriction(pr+"\u2192", board, rows, cols+1)
		paths = append(paths, horizPaths...)
	}
	if rows < len(board)-1 {
		vertPaths := pathsWithRestriction(pr+"\u2193", board, rows+1, cols)
		paths = append(paths, vertPaths...)
	}
	return paths
}

func allPathsWithRestriction(pr string, board [][]bool, rows, cols int) []string {
	paths := make([]string, 0)
	if rows == len(board)-1 && cols == len(board[0])-1 {
		return []string{pr}
	}

	if !board[rows][cols] {
		return paths
	}
	// block already considered for this path
	board[rows][cols] = false

	// down-right diag
	if rows < len(board)-1 && cols < len(board[0])-1 {
		downright := allPathsWithRestriction(pr+"\u2198", board, rows+1, cols+1)
		paths = append(paths, downright...)
	}
	// down-left diag
	if rows < len(board)-1 && cols > 0 {
		downleft := allPathsWithRestriction(pr+"\u2199", board, rows+1, cols-1)
		paths = append(paths, downleft...)
	}
	// up-left diag
	if rows > 0 && cols > 0 {
		upleft := allPathsWithRestriction(pr+"\u2196", board, rows-1, cols-1)
		paths = append(paths, upleft...)
	}
	// up-right diag
	if rows > 0 && cols < len(board[0])-1 {
		upright := allPathsWithRestriction(pr+"\u2197", board, rows-1, cols+1)
		paths = append(paths, upright...)
	}

	// left
	if cols > 0 {
		leftPaths := allPathsWithRestriction(pr+"\u2190", board, rows, cols-1)
		paths = append(paths, leftPaths...)
	}
	// right
	if cols < len(board[0])-1 {
		rightPaths := allPathsWithRestriction(pr+"\u2192", board, rows, cols+1)
		paths = append(paths, rightPaths...)
	}
	// up
	if rows > 0 {
		upPaths := allPathsWithRestriction(pr+"\u2191", board, rows-1, cols)
		paths = append(paths, upPaths...)
	}
	// down
	if rows < len(board)-1 {
		downPath := allPathsWithRestriction(pr+"\u2193", board, rows+1, cols)
		paths = append(paths, downPath...)
	}

	// backtracking to the original state
	board[rows][cols] = true
	return paths
}
