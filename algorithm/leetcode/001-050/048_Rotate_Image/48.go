package _48_Rotate_Image

func rotate(matrix [][]int) {
	top, down := 0, len(matrix)-1
	for top < down {
		matrix[top], matrix[down] = matrix[down], matrix[top]
		top++
		down--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
