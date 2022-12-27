package _74_Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	} else if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0] == target
	}
	row := len(matrix)
	column := len(matrix[0])
	max := row * column
	l, r := 0, max-1
	for l+1 < r {
		mid := (r-l)/2 + l
		num := matrix[mid/column][mid%column]
		if num == target {
			return true
		} else if num > target {
			r = mid
		} else {
			l = mid
		}
	}
	//max == 2
	if matrix[l/column][l%column] == target {
		return true
	} else if matrix[r/column][r%column] == target {
		return true
	}
	return false
}
