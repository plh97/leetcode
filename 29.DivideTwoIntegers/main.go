package divide

import "math"

func divide(dividend int, divisor int) int {
	res := 0
	if dividend < 0 && divisor < 0 {
		// -1 -1
		for dividend <= divisor {
			dividend -= divisor
			res++
		}
	} else if dividend > 0 && divisor > 0 {
		// 1 1
		for dividend >= divisor {
			dividend -= divisor
			res++
		}
	} else if dividend < 0 && divisor > 0 {
		// -1 1
		for -dividend >= divisor {
			dividend += divisor
			res--
		}
	} else if dividend > 0 && divisor < 0 {
		// 1 -1
		for dividend >= -divisor {
			dividend += divisor
			res--
		}
	}
	if res >= math.MaxInt32 {
		return math.MaxInt32
	} else if res < math.MinInt32 {
		return math.MinInt32
	}
	return res
}
