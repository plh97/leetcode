package thirdMax

import (
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	// 去重
	for i := 0; i < len(nums); i++ {
		lastIndexOf(nums, 888888)
		if i != lastIndexOf(nums, nums[i]) {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	if len(nums) > 2 {
		return nums[len(nums)-3]
	} else {
		return nums[len(nums)-1]
	}
}

func lastIndexOf(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
