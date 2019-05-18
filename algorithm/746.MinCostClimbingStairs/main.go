package minCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
