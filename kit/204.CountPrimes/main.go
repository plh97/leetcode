package countPrimes

func countPrimes(n int) int {
	c := 0
	dp := make([]bool, n)
	for i := 2; i < n; i++ {
		// n is 素数
		if !dp[i] {
			c++
			for j := i; j < n; j += i {
				dp[j] = true
			}
		}
	}
	return c
}
