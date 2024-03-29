package _40_Search_a_2D_Matrix_II

func searchMatrix(matrix [][]int, target int) bool {
	row, column := 0, len(matrix[0])-1
	for row < len(matrix) && column >= 0 {
		if target == matrix[row][column] {
			return true
		} else if target > matrix[row][column] {
			row++
		} else if target < matrix[row][column] {
			column--
		}
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		}
	}
	return false
}
