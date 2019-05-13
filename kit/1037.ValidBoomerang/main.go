package isBoomerang

import "math"

func isBoomerang(points [][]int) bool {
	c := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
	}
	for i := range c {
		if points[c[i][0]][0] == points[c[i][1]][0] &&
			points[c[i][0]][1] == points[c[i][1]][1] {
			return false
		}
	}
	dx01 := points[1][0] - points[0][0]
	dy01 := points[1][1] - points[0][1]
	dx02 := points[2][0] - points[0][0]
	dy02 := points[2][1] - points[0][1]
	xxx01 := 0.0
	xxx02 := 0.0
	if dy01 == 0 {
		xxx01 = math.MaxFloat64
	} else {
		xxx01 = float64(dx01) / float64(dy01)
	}
	if dy02 == 0 {
		xxx02 = math.MaxFloat64
	} else {
		xxx02 = float64(dx02) / float64(dy02)
	}
	if xxx01 == xxx02 {
		return false
	}
	return true
}
