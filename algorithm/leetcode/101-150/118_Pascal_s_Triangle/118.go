package _18_Pascal_s_Triangle

func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		var levels []int
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				levels = append(levels, 1)
			} else {
				levels = append(levels, result[i-1][j]+result[i-1][j-1])
			}
		}
		result = append(result, levels)
	}
	return result
}
