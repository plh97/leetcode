package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序函数
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	for i := 1; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := 0, len(nums)-1
		for l < r {
			sum := nums[l] + nums[i] + nums[r]
			fmt.Println(sum, l, r)
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{nums[l], nums[i], nums[r]})
				// 此处为了防止重复, `0,0,0,1,0,0,1`这种情况
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	switch {
	case nums[l] == nums[l+1]:
		l++
	case nums[r] == nums[r-1]:
		r--
	default:
		l++
		r--
	}
	return l, r
}

// func indexOf(element int, data []int) int {
// 	for k, v := range data {
// 		if element == v {
// 			return k
// 		}
// 	}
// 	return -1 //not found.
// }

func main() {
	fmt.Println(threeSum([]int{4, 4, 3, -5, 0, 0, 0, -2, 3, -5, -5, 0}))
}
