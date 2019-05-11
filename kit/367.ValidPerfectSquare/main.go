package isPerfectSquare

import "math"

func isPerfectSquare(num int) bool {
	if int(math.Sqrt(float64(num)))*int(math.Sqrt(float64(num))) == num {
		return true
	}
	return false
}
