package dominantIndex

func dominantIndex(nums []int) int {
	max1 := -1 << 63
	max2 := -1 << 63
	Map := make(map[int]int, len(nums))
	for i, e := range nums {
		Map[e] = i
		if e > max1 {
			max2 = max1
			max1 = e
		} else if e > max2 {
			max2 = e
		}
	}
	if max1 >= max2*2 {
		return Map[max1]
	}
	return -1
}
