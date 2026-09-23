package main

import (
	"fmt"
	"strconv"
)

// my solution
func isValidSudoku(board [][]byte) bool {
	map3x3 := make(map[string]map[byte]bool)

	for i := range board {
		mapRow := make(map[byte]bool)
		mapCol := make(map[byte]bool)

		for j := range board[i] {
			row, col := board[i][j], board[j][i]

			if row != '.' {
				if _, ok := mapRow[row]; ok {
					return false
				} else {
					mapRow[row] = true
				}

				ij := strconv.Itoa(i/3) + strconv.Itoa(j/3)

				block3x3, ok := map3x3[ij]
				if !ok {
					block3x3 = make(map[byte]bool)
					map3x3[ij] = block3x3
				}

				if _, ok := block3x3[row]; ok {
					return false
				} else {
					block3x3[row] = true
				}
			}

			if col != '.' {
				if _, ok := mapCol[col]; ok {
					return false
				} else {
					mapCol[col] = true
				}
			}
		}
	}

	return true
}

func isValidSudokuOptimized(board [][]byte) bool {
	var blocksBits [9]uint16
	rowBits, colBits := uint16(0), uint16(0)

	for i := range board {
		for j := range board[i] {
			vRow, vCol := board[i][j]-'0'-1, board[j][i]-'0'-1 // < 0 if '.'

			if vRow >= 0 {
				row := uint16(1 << vRow)
				block := (i/3)*3 + (j / 3)

				if rowBits&row != 0 || blocksBits[block]&row != 0 {
					return false
				}

				rowBits |= row
				blocksBits[block] |= row
			}

			if vCol >= 0 {
				col := uint16(1 << vCol)

				if colBits&col != 0 {
					return false
				}

				colBits |= col
			}
		}

		rowBits, colBits = 0, 0
	}

	return true
}

func main() {
	fmt.Println(isValidSudokuOptimized([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}
