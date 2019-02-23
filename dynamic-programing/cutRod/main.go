package cutRod

import (
	"fmt"
	"math"
)

var i int
var dp map[int]int = make(map[int]int)

// p -> 价格参数集合
func cutRodDp(p [8]int, n int) int {
	i++
	fmt.Println("-次数", i, n, dp)
	if res, ok := dp[n]; ok {
		return res
	}
	if n == 0 {
		return 0
	}
	q := math.MinInt8
	for i := 1; i <= n; i++ {
		if q < (p[i] + cutRodDp(p, n-i)) {
			q = p[i] + cutRodDp(p, n-i)
		}
	}
	dp[n] = q
	return q
}




func cutRod(p [8]int, n int) int {
	i++
	fmt.Println("次数", i, n)
	if n == 0 {
		return 0
	}
	q := math.MinInt8
	for i := 1; i <= n; i++ {
		if q < (p[i] + cutRod(p, n-i)) {
			q = p[i] + cutRod(p, n-i)
		}
	}
	return q
}
