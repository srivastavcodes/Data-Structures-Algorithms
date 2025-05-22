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

	// Base case: reached destination
	if rows == len(board)-1 && cols == len(board[0])-1 {
		return []string{pr}
	}
	// Skip if cell is blocked
	if !board[rows][cols] {
		return paths
	}
	// Mark current cell as visited
	board[rows][cols] = false

	paths = append(paths, handleCardinalPaths(pr, board, rows, cols)...)

	paths = append(paths, handleDiagonalPaths(pr, board, rows, cols)...)

	// Backtracking: mark cell as unvisited again
	board[rows][cols] = true

	return paths
}

// Handle cardinal directions: up, down, left, right
func handleCardinalPaths(pr string, board [][]bool, rows, cols int) []string {
	paths := make([]string, 0)

	// Down movement
	if rows < len(board)-1 {
		downPaths := allPathsWithRestriction(pr+"\u2193", board, rows+1, cols)
		paths = append(paths, downPaths...)
	}
	// Right movement
	if cols < len(board[0])-1 {
		rightPaths := allPathsWithRestriction(pr+"\u2192", board, rows, cols+1)
		paths = append(paths, rightPaths...)
	}
	// Up movement
	if rows > 0 {
		upPaths := allPathsWithRestriction(pr+"\u2191", board, rows-1, cols)
		paths = append(paths, upPaths...)
	}
	// Left movement
	if cols > 0 {
		leftPaths := allPathsWithRestriction(pr+"\u2190", board, rows, cols-1)
		paths = append(paths, leftPaths...)
	}
	return paths
}

// Handle diagonal directions: up-left, upright, down-left, down-right
func handleDiagonalPaths(pr string, board [][]bool, rows, cols int) []string {
	paths := make([]string, 0)

	// Down-right diagonal
	if rows < len(board)-1 && cols < len(board[0])-1 {
		downRightPaths := allPathsWithRestriction(pr+"\u2198", board, rows+1, cols+1)
		paths = append(paths, downRightPaths...)
	}
	// Up-left diagonal
	if rows > 0 && cols > 0 {
		upLeftPaths := allPathsWithRestriction(pr+"\u2196", board, rows-1, cols-1)
		paths = append(paths, upLeftPaths...)
	}
	// Down-left diagonal
	if rows < len(board)-1 && cols > 0 {
		downLeftPaths := allPathsWithRestriction(pr+"\u2199", board, rows+1, cols-1)
		paths = append(paths, downLeftPaths...)
	}
	// Up-right diagonal
	if rows > 0 && cols < len(board[0])-1 {
		upRightPaths := allPathsWithRestriction(pr+"\u2197", board, rows-1, cols+1)
		paths = append(paths, upRightPaths...)
	}
	return paths
}
