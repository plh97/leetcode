package missingNumber

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if len(nums) == 1 {
			//第一个数字
			if nums[0] == 0 {
				return 1
			} else if nums[0] == 1 {
				return 0
			}
		} else if i < len(nums)-1 && nums[i+1]-nums[i] != 1 {
			// 后面的数字
			return i + 1
		}
	}
	if nums[0] < 1 {
		return len(nums)
	}
	return 0
}
