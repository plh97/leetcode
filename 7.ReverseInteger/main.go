package convert

import (
	"math"
)

func reverse(x int) int {
	sign := 1
	res := 0
	if x < 0 {
		sign = -1
		x = -1 * x
	}
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}
	res = sign * res
	if res > math.MaxInt32 || res < -math.MaxInt32 {
		return 0
	}
	return res
}
