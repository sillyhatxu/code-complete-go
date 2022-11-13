package _30_Surrounded_Regions

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	lastRow, lastColumn := len(board)-1, len(board[0])-1
	for i := 0; i <= lastColumn; i++ {
		if board[0][i] == 'O' {
			helper(&board, 0, i)
		}
		if board[lastRow][i] == 'O' {
			helper(&board, lastRow, i)
		}
	}
	for i := 0; i <= lastRow; i++ {
		if board[i][0] == 'O' {
			helper(&board, i, 0)
		}
		if board[i][lastColumn] == 'O' {
			helper(&board, i, lastColumn)
		}
	}
	for i := 0; i <= lastRow; i++ {
		for j := 0; j <= lastColumn; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}
}

func helper(board *[][]byte, x, y int) {
	if x < 0 || x > len(*board)-1 || y < 0 || y > len((*board)[0])-1 || (*board)[x][y] != 'O' {
		return
	}
	(*board)[x][y] = '-'
	helper(board, x-1, y)
	helper(board, x+1, y)
	helper(board, x, y-1)
	helper(board, x, y+1)
}
