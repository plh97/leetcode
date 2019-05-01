package findMaxAverage

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	res := -math.MaxFloat64

	for i := 0; i <= len(nums)-k; i++ {
		temp := nums[i : i+k]
		totalCount := 0.0
		for j := 0; j < k; j++ {
			totalCount += float64(temp[j])
		}
		res = math.Max(res, float64(totalCount / float64(k)))
	}

	return res
}
