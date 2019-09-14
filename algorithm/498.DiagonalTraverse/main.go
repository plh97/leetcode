package findDiagonalOrder

func findDiagonalOrder(matrix [][]int) []int {
	res := []int{}
	x := 0
	y := 0
	direct := true
	for x < len(matrix[0]) && y < len(matrix) {
		res = append(res, matrix[y][x])
		if direct {
			// true 	- |
			// true 	↗ |
			if y == 0 || x == len(matrix[0])-1 {
				// react bounder
				direct = !direct
				if x == len(matrix[0])-1 {
					y++
				} else {
					x++
				}
			} else {
				y--
				x++
			}
		} else {
			// false	| ↙️
			// false	|--
			if x == 0 || y == len(matrix)-1 {
				// react bounder
				direct = !direct
				if y == len(matrix)-1 {
					x++
				} else {
					y++
				}
			} else {
				y++
				x--
			}
		}
	}
	return res
}
