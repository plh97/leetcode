package spiralOrder

var hasSeen map[int]bool

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	hasSeen = map[int]bool{}
	lenx := len(matrix[0])
	leny := len(matrix)
	res := []int{}
	x, y := 0, 0
	direct := 0
	for i := 0; i < lenx*leny; i++ {

		res = append(res, matrix[y][x])
		hasSeen[matrix[y][x]] = true
		if isBound(matrix, x, y, direct) {
			direct++
		}
		switch direct % 4 {
		case 0:
			// right 	0
			x++
		case 1:
			// down 	1
			y++
		case 2:
			// left 	2
			x--
		case 3:
			// up 		3
			y--
		}
	}
	return res
}

func isBound(matrix [][]int, x, y, direct int) bool {
	switch direct % 4 {
	case 0:
		x++
	case 1:
		y++
	case 2:
		x--
	case 3:
		y--
	}
	if x < 0 ||
		y < 0 ||
		y > len(matrix)-1 ||
		x > len(matrix[0])-1 {
		return true
	}
	if _, ok := hasSeen[matrix[y][x]]; ok {
		return true
	}
	return false
}
