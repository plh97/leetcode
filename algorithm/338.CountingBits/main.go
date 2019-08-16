package countBits

func countBits(num int) []int {
	res := []int{}
	for i := 0; i <= num; i++ {
		pushNum := 0
		j := i
		for j > 0 {
			if j%2 == 1 {
				pushNum++
			}
			j /= 2
		}
		res = append(res, pushNum)
	}
	return res
}
