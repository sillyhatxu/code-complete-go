package _37_Sudoku_Solver

import "fmt"

//TODO
func solveSudoku(board [][]byte) {
	var rows [][]bool
	var columns [][]bool
	var blocks [][]bool
	initial(board, &rows, &columns, &blocks)
	dfs(&board, &rows, &columns, &blocks, 0, 0)
	fmt.Println(board)
}

func dfs(board *[][]byte, rows *[][]bool, columns *[][]bool, blocks *[][]bool, i, j int) {
	number := (*board)[i][j]
	if number != '.' {
		nextI, nextJ := nextCell(i, j)
		if nextI == -1 {
			return
		}
		dfs(board, rows, columns, blocks, nextI, nextJ)
		return
	}
	for cell := '1'; cell <= '9'; cell++ {

	}

}

func nextCell(i, j int) (int, int) {
	if i == 8 && j == 8 {
		return -1, -1
	} else if j == 8 {
		return i + 1, 0
	} else {
		return i, j + 1
	}
}

func initial(board [][]byte, rows *[][]bool, columns *[][]bool, blocks *[][]bool) {
	*rows = make([][]bool, 9)
	*columns = make([][]bool, 9)
	*blocks = make([][]bool, 9)
	for i := 0; i < 9; i++ {
		(*rows)[i] = make([]bool, 9)
		(*columns)[i] = make([]bool, 9)
		(*blocks)[i] = make([]bool, 9)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				number := board[i][j] - '1'
				(*rows)[i][number] = true
				(*columns)[j][number] = true
				(*blocks)[i/3*3+j/3][number] = true
			}
		}
	}
}

//var rows [][]bool
//var columns [][]bool
//var blocks [][]bool
