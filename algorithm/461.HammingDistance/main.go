package hammingDistance

func hammingDistance(x int, y int) int {

	res := 0
	for diff := x ^ y; diff > 0; diff /= 2 {
		if diff%2 == 1 {
			res++
		}
	}
	return res
}
