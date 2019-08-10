package maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	leni := len(grid)
	lenj := len(grid[0])
	for i := 0; i < leni; i++ {
		for j := 0; j < lenj; j++ {
			e := grid[i][j]
			if e == 1 {
				ress := helper(&grid, i, j)
				if ress > res {
					res = ress
				}
			}
		}
	}
	return res
}

// 基于当前节点扩散搜索..
func helper(grid *[][]int, i, j int) int {
	if isPointInner(*grid, i, j) && (*grid)[i][j] == 1 {
		(*grid)[i][j] = 0
		return 1 + helper(grid, i+1, j) + helper(grid, i-1, j) + helper(grid, i, j+1) + helper(grid, i, j-1)
	} else {
		return 0
	}
}

// 范围内部 && val==0 : true
func isPointInner(grid [][]int, i int, j int) bool {
	if i < 0 ||
		j < 0 ||
		i >= len(grid) ||
		j >= len(grid[0]) {
		return false
	}
	return true
}
