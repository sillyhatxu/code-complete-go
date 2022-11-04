package _54_Spiral_Matrix

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	rowStart, rowEnd, columnStart, columnEnd := 0, len(matrix)-1, 0, len(matrix[0])-1
	for rowStart <= rowEnd && columnStart <= columnEnd {
		for columnIndex := columnStart; columnIndex <= columnEnd; columnIndex++ {
			result = append(result, matrix[rowStart][columnIndex])
		}
		rowStart++
		if rowStart > rowEnd {
			continue
		}
		for rowIndex := rowStart; rowIndex <= rowEnd; rowIndex++ {
			result = append(result, matrix[rowIndex][columnEnd])
		}
		columnEnd--
		if columnStart > columnEnd {
			continue
		}
		for columnIndex := columnEnd; columnIndex >= columnStart; columnIndex-- {
			result = append(result, matrix[rowEnd][columnIndex])
		}
		rowEnd--
		if rowStart > rowEnd {
			continue
		}
		for rowIndex := rowEnd; rowIndex >= rowStart; rowIndex-- {
			result = append(result, matrix[rowIndex][columnStart])
		}
		columnStart++
	}
	return result
}
