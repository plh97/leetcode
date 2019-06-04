package maxProfit

func maxProfit(prices []int) int {
	res := 0
	for i, e := range prices {
		if i == len(prices)-1 {
			continue
		}
		res += max(prices[i+1]-e, 0)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}