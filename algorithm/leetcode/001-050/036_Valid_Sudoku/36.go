package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
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

func main() {
	fmt.Println(0 / 3)
	fmt.Println(1 / 3)
	fmt.Println(2 / 3)
	fmt.Println(3 / 3)
	fmt.Println(4 / 3)
	fmt.Println(5 / 3)
	fmt.Println(6 / 3)
	fmt.Println(7 / 3)
	fmt.Println(8 / 3)

	//isValidSudoku(nil)
}
