package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	solutions := make([]string, 0)
	solveEightQueens(0, make([]int, 8), &solutions)
	return strings.Join(solutions, "\n")
}

func solveEightQueens(row int, board []int, solutions *[]string) {
	if row == 8 {
			solution := ""
			for _, col := range board {
					solution += strconv.Itoa(col + 1)
			}
			*solutions = append(*solutions, solution)
			return
	}

	for col := 0; col < 8; col++ {
			if isSafe(row, col, board) {
					board[row] = col
					solveEightQueens(row+1, board, solutions)
					board[row] = -1
			}
	}
}

func isSafe(row, col int, board []int) bool {
	for i := 0; i < row; i++ {
			if board[i] == col || board[i] == col-row+i || board[i] == col+row-i {
					return false
			}
	}
	return true
}