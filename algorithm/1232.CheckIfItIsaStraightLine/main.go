package checkstraightline

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 0 || len(coordinates) == 1 {
		return true
	}
	var dx float64 = float64(coordinates[1][0] - coordinates[0][0])
	var dy float64 = float64(coordinates[1][1] - coordinates[0][1])
	var dd float64 = dy / dx
	d0 := coordinates[0]
	for i, coord := range coordinates {
		if i == 0 || i == 1 {
			continue
		} else if float64(coord[1]-d0[1])/float64(coord[0]-d0[0]) != dd {
			return false
		}
	}
	return true
}
