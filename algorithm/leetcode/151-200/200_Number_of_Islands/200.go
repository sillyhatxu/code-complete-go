package _00_Number_of_Islands

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				drawingIsland(&grid, i, j)
			}
		}
	}
	return res
}

func drawingIsland(grid *[][]byte, i, j int) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) || (*grid)[i][j] != '1' {
		return
	}
	(*grid)[i][j] = '2'
	drawingIsland(grid, i-1, j)
	drawingIsland(grid, i+1, j)
	drawingIsland(grid, i, j-1)
	drawingIsland(grid, i, j+1)
}
