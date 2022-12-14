package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	validMap := make(map[string]struct{})
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			number := board[row][col]
			if number == '.' {
				continue
			}
			keyRow := fmt.Sprintf("r[%d]%d", row, number)
			keyCol := fmt.Sprintf("c[%d]%d", col, number)
			keyBlock := fmt.Sprintf("b[%d]%d-%d", number, row/3, col/3)
			if _, ok := validMap[keyRow]; ok {
				return false
			} else if _, ok := validMap[keyCol]; ok {
				return false
			} else if _, ok := validMap[keyBlock]; ok {
				return false
			}
			validMap[keyRow] = struct{}{}
			validMap[keyCol] = struct{}{}
			validMap[keyBlock] = struct{}{}
		}
	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
	rowMap := [9][9]bool{}
	columnMap := [9][9]bool{}
	gridMap := [9][9]bool{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cell := string(board[i][j])
			if cell == "." {
				continue
			}
			val, _ := strconv.Atoi(cell)
			val--
			gridIndex := i/3*3 + j/3
			if rowMap[i][val] || columnMap[j][val] || gridMap[gridIndex][val] {
				return false
			}
			rowMap[i][val] = true
			columnMap[j][val] = true
			gridMap[gridIndex][val] = true
		}
	}
	return true
}
