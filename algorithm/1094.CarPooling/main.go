package carPooling

func carPooling(trips [][]int, capacity int) bool {
	for i, max := 0, 1; i < max; i++ {
		num := 0
		for _, trip := range trips {
			if trip[1] <= i && trip[2] > i {
				num += trip[0]
			}
			if max < trip[2] {
				max = trip[2]
			}
		}
		if capacity < num {
			return false
		}
	}
	return true
}
