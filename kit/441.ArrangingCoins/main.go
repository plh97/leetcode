package arrangeCoins

func arrangeCoins(n int) int {
	res := 0
	for i := 1; n >= i; i++ {
		n -= i
		res = i
	}
	return res
}
