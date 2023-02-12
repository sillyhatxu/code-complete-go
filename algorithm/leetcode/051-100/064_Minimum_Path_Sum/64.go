package _64_Minimum_Path_Sum

func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	m, n := len(grid), len(grid[0])
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func minPathSum1(grid [][]int) int {
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
