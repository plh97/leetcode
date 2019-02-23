package fibDp

var dp map[int]int = make(map[int]int)

// 动态规划版本的斐波那契数列 可一秒内计算到 fib(3521155), 调用7M次函数进行递归 这是c++的性能
func fibDp(n int) int {
	if res, ok := dp[n]; ok {
		return res
	}
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}
	res := fibDp(n-1) + fibDp(n-2)
	dp[n] = res
	return res
}

// 直接递归, 没有动态规划最大 fib(43)
func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return fib(n-1) + fib(n-2)
}
