package _94_Rotting_Oranges

func orangesRotting(grid [][]int) int {
	fresh, rotten, m, n := 1, 2, len(grid), len(grid[0])
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} //下一步的方向，[下，上，右，左]
	var rottenQueue [][2]int
	freshCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case fresh:
				freshCount++
			case rotten:
				rottenQueue = append(rottenQueue, [2]int{i, j})
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	res := 0
	for len(rottenQueue) > 0 {
		for _, ri := range rottenQueue {
			for _, dir := range directions {
				x, y := ri[0]+dir[0], ri[1]+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == fresh {
					grid[x][y] = rotten
					rottenQueue = append(rottenQueue, [2]int{x, y})
					freshCount--
				}
			}
			rottenQueue = rottenQueue[1:]
		}
		res++
		if freshCount == 0 {
			return res
		}
	}
	return -1
}
