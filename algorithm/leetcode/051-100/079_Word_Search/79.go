package _79_Word_Search

/**
Time: O(4^1), l = len(word)
total: O(m*n*4^1)

Space: O(m*n+1)
*/

/**
DFS without extra space
*/
//func exist(board [][]byte, word string) bool {
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[i]); j++ {
//			if board[i][j] == word[0] {
//				result := dfs(board, word, i, j, 0)
//				if result {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
//
//func dfs(board [][]byte, word string, rowIndex, colIndex int, depth int) bool {
//	if rowIndex < 0 || colIndex < 0|| rowIndex >= len(board) || colIndex >= len(board[0]) || depth >= len(word) || word[depth : depth+1] != string(board[rowIndex][colIndex]) {
//		return false
//	}
//	if depth == len(word) -1{
//		return true
//	}
//	current := board[rowIndex][colIndex]
//	board[rowIndex][colIndex] = '-'
//	result := dfs(board, word, rowIndex, colIndex+1, depth+1) ||
//		dfs(board, word, rowIndex, colIndex-1, depth+1) ||
//		dfs(board, word, rowIndex+1, colIndex, depth+1) ||
//		dfs(board, word, rowIndex-1, colIndex, depth+1)
//	board[rowIndex][colIndex] = current
//	return result
//}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				result := dfs(board, word, i, j, 0)
				if result {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, rowIndex, colIndex int, depth int) bool {
	if board[rowIndex][colIndex] != word[depth] || board[rowIndex][colIndex] == '-' {
		return false
	}
	if depth == len(word)-1 {
		return true
	}
	current := board[rowIndex][colIndex]
	board[rowIndex][colIndex] = '-'
	depth++
	if rowIndex > 0 && dfs(board, word, rowIndex-1, colIndex, depth) {
		return true
	}
	if rowIndex < len(board)-1 && dfs(board, word, rowIndex+1, colIndex, depth) {
		return true
	}
	if colIndex > 0 && dfs(board, word, rowIndex, colIndex-1, depth) {
		return true
	}
	if colIndex < len(board[0])-1 && dfs(board, word, rowIndex, colIndex+1, depth) {
		return true
	}
	board[rowIndex][colIndex] = current
	return false
}
