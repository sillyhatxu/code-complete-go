package _62_Unique_Paths

func uniquePaths(m int, n int) int {
	row := make([]int, n)
	row[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == 0 {
				row[j] = 1
				continue
			}
			row[j] += row[j-1]
		}
	}
	return row[n-1]
}

func uniquePaths1(m int, n int) int {
	grid := make(map[int]map[int]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make(map[int]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
				continue
			}
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 1; j < n; j++ {
			if i == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
