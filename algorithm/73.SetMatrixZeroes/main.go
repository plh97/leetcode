package setZeroes

func setZeroes(matrix [][]int) [][]int {
	xx := []int{}
	yy := []int{}
	for y, cols := range matrix {
		for x, e := range cols {
			if e == 0 {
				xx = append(xx, x)
				yy = append(yy, y)
			}
		}
	}
	for y, cols := range matrix {
		for x := range cols {
			if indexOf(xx, x) > -1 || indexOf(yy, y) > -1 {
				matrix[y][x] = 0
			}
		}
	}
	return matrix
}

func indexOf(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
