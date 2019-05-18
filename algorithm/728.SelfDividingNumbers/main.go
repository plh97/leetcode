package selfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		if isSelfDevide(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDevide(n int) bool {
	nn := n
	for nn > 0 {
		p := nn % 10
		if p == 0 || n%p != 0 {
			return false
		}
		nn /= 10
	}
	return true
}
