package islandPerimeter

func islandPerimeter(grid [][]int) int {
	res := 0
	// row to for eachc
	for i, col := range grid {
		for j, e := range col {
			if e == 1 {
				res += 4
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}
