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
