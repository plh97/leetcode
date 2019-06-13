package maxIncreaseKeepingSkyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	res := 0
	colMax := []int{}
	rowMax := []int{}

	// 寻找横向最大值
	for i := 0; i < len(grid); i++ {
		max := grid[i][0]
		for j := 1; j < len(grid[i]); j++ {
			max = maxFunc(grid[i][j], max)
		}
		colMax = append(colMax, max)
	}

	// 寻找纵向最大值
	for i := 0; i < len(grid[0]); i++ {
		max := grid[0][i]
		for j := 1; j < len(grid); j++ {
			max = maxFunc(grid[j][i], max)
		}
		rowMax = append(rowMax, max)
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			e := grid[i][j]
			res += minFunc(rowMax[j], colMax[i]) - e
		}
	}

	return res
}
func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minFunc(a, b int) int {
	if a < b {
		return a
	}
	return b
}
