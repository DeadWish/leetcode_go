func islandPerimeter(grid [][]int) int {
	islands := 0
	repeated := 0
	xlen := len(grid)
	ylen := len(grid[0])
	for x, xline := range grid {
		for y, _ := range xline {
			if grid[x][y] == 1 {
				islands++
				if x < xlen-1 && grid[x+1][y] == 1 {
					repeated++
				}
				if y < ylen-1 && grid[x][y+1] == 1 {
					repeated++
				}
			}
		}
	}

	return islands*4 - repeated*2
}