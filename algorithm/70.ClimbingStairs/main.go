package climbStairs

// dp

var dp map[int]int = make(map[int]int, 45)

func climbStairs(n int) int {
	dp[1] = 1
	dp[2] = 2
	return helper(n)
}
func helper(n int) int {
	if e, ok := dp[n]; ok {
		return e
	}
	dp[n] = helper(n-1) + helper(n-2)
	return dp[n]
}
