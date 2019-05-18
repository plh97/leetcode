package maxProfit

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i+1; j < len(prices); j++ {
			temp := prices[j] - prices[i]
			if res < temp {
				res = temp
			}
		}
	}
	return res
}
