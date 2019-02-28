package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序函数
	res := [][]int{}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[i] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 此处为了防止重复, `0,0,0,1,0,0,1`这种情况
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}
