package _89_Game_of_Life

func gameOfLife(board [][]int) {
	rowIndex, colIndex := len(board), len(board[0])
	for i := 0; i < rowIndex; i++ {
		for j := 0; j < colIndex; j++ {
			neighbor := neighborCount(board, i, j)
			if board[i][j] == 1 {
				if neighbor == 2 || neighbor == 3 {
					board[i][j] += 2
				}
			} else if board[i][j] == 0 {
				if neighbor == 3 {
					board[i][j] += 2
				}
			}
		}
	}
	for i := 0; i < rowIndex; i++ {
		for j := 0; j < colIndex; j++ {
			board[i][j] >>= 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func neighborCount(board [][]int, rowIndex, colIndex int) int {
	count := 0
	for i := max(0, rowIndex-1); i <= min(rowIndex+1, len(board)-1); i++ {
		for j := max(0, colIndex-1); j <= min(colIndex+1, len(board[0])-1); j++ {
			if i == rowIndex && j == colIndex {
				continue
			}
			if board[i][j]&1 == 1 {
				count++
			}
		}
	}
	return count
}
