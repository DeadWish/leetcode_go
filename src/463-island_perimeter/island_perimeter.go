func islandPerimeter(grid [][]int) int {
	sx, sy := 0, 0
	flag := false
	for x, sv := range grid {
		if flag {
			break
		}
		for y, b := range sv {
			if flag {
				break
			}
			if b == 1 {
				sx, sy = x, y
				flag = true
			}
		}
	}

	return cal(sx, sy, &grid)
}

func cal(x int, y int, grid *[][]int) int {
	if (*grid)[x][y] == 2 {
		return 0
	}
	if (*grid)[x][y] != 1 {
		return 1
	}

	(*grid)[x][y] = 2
	perimeter := 0
	if tempx := x - 1; tempx >= 0 {
		perimeter += cal(tempx, y, grid)
	} else {
		perimeter++
	}
	if tempx := x + 1; tempx <= len(*grid)-1 {
		perimeter += cal(tempx, y, grid)
	} else {
		perimeter++
	}
	if tempy := y - 1; tempy >= 0 {
		perimeter += cal(x, tempy, grid)
	} else {
		perimeter++
	}
	if tempy := y + 1; tempy <= len((*grid)[0])-1 {
		perimeter += cal(x, tempy, grid)
	} else {
		perimeter++
	}

	return perimeter
}