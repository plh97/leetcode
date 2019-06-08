package rob

var dp map[int]int

// [1,2,3,4]
func rob(nums []int) int {
	dp = make(map[int]int, len(nums)+1)
	dp[0] = 0
	return helper(nums)
}
func helper(nums []int) int {
	if len(nums) == 1 {
		dp[1] = nums[0]
	}
	if len(nums) == 2 {
		dp[2] = max(nums[0], nums[1])
	}
	if e, ok := dp[len(nums)]; ok {
		return e
	}
	// 分两个情况, 抢劫第一家, 不抢劫第一家
	dp[len(nums)] = max(helper(nums[1:]), nums[0]+helper(nums[2:]))
	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
