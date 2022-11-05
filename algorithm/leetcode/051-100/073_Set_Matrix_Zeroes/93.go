package _73_Set_Matrix_Zeroes

func setZeroes(matrix [][]int) {
	zeroInRow, zeroInColumn := false, false
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				if row == 0 {
					zeroInRow = true
				}
				if col == 0 {
					zeroInColumn = true
				}
				matrix[0][col], matrix[row][0] = 0, 0
			}
		}
	}
	for col := 1; col < len(matrix[0]); col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < len(matrix); row++ {
				matrix[row][col] = 0
			}
		}
	}
	for row := 1; row < len(matrix); row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < len(matrix[0]); col++ {
				matrix[row][col] = 0
			}
		}
	}
	if zeroInRow {
		for col := 0; col < len(matrix[0]); col++ {
			matrix[0][col] = 0
		}
	}
	if zeroInColumn {
		for row := 0; row < len(matrix); row++ {
			matrix[row][0] = 0
		}
	}
}
