package _64_Minimum_Path_Sum

func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	row := len(grid)
	column := len(grid[0])
	for c := 1; c < column; c++ {
		grid[0][c] += grid[0][c-1]
	}
	for r := 1; r < row; r++ {
		grid[r][0] += grid[r-1][0]
	}
	for r := 1; r < row; r++ {
		for c := 1; c < column; c++ {
			grid[r][c] += min(grid[r][c-1], grid[r-1][c])
		}
	}
	return grid[row-1][column-1]
}
