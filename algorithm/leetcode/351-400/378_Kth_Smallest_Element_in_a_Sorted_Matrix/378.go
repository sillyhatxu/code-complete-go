package _78_Kth_Smallest_Element_in_a_Sorted_Matrix

func kthSmallest(matrix [][]int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	row, col := 0, 0
	for k > 0 {
		k--
		if k == 0 {
			return matrix[row][col]
		}
		if matrix[min(row+1, len(matrix)-1)][col] >= matrix[row][min(col+1, len(matrix[0])-1)] {
			if col == len(matrix[0])-1 {
				row++
			} else {
				col++
			}
		} else if matrix[min(row+1, len(matrix)-1)][col] < matrix[row][min(col+1, len(matrix[0])-1)] {
			if row == len(matrix)-1 {
				col++
			} else {
				row++
			}
		}
	}
	//queue := make([][]int, len(matrix), len(matrix))
	//for i := 0; i < len(matrix); i++ {
	//	queue[matrix[i][0]] = []int{i, 0}
	//}
	//for k > 0 {
	//	if k == 0{
	//
	//	}
	//	k--
	//}
	return -1
}
