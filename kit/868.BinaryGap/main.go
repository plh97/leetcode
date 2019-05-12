package binaryGap

import (
	"math"
)

func binaryGap(N int) int {
	res := 0.0
	currentIndex := -1
	for i := 0; N > 0; i++ {
		currentDigit := N % 2
		if currentDigit == 1 {
			if currentIndex > -1 {
				res = math.Max(res, float64(i-currentIndex))
			}
			currentIndex = i
		}
		N /= 2
	}
	return int(res)
}
