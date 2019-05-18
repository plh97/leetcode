package judgeSquareSum

import "math"

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		temp := c - i*i
		float_c := int(math.Sqrt(float64(temp)))
		if float_c*float_c+i*i == c {
			return true
		}
	}
	return false
}
