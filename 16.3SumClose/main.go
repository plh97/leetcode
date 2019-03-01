package threesum

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	sort.Ints(nums) // 排序函数
	res := 9999
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if math.Abs(float64(res-target)) > math.Abs(float64(nums[i]+nums[j]+nums[k]-target)) {
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return res
}
