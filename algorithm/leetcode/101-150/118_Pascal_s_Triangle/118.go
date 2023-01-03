package _18_Pascal_s_Triangle

func generate(numRows int) [][]int {
	var res [][]int
	for i := 1; i <= numRows; i++ {
		temp := make([]int, i, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				temp[j] = 1
			} else {
				temp[j] = res[i-2][j-1] + res[i-2][j]
			}
		}
		res = append(res, temp)
	}
	return res
}

func generate1(numRows int) [][]int {
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
