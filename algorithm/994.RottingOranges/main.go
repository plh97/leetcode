package orangesRotting

import "fmt"

func orangesRotting(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			e := grid[i][j]
			if e == 2 {
				fmt.Println(i, j, grid[i][j])
			}
		}
	}
	return res
}
