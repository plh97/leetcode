package majorityElement

import "sort"

func majorityElement(nums []int) int {

	// ints := []int{7, 2, 4}
	sort.Ints(nums)
	// res := sort.Ints([]int{1, 2, 3, 45, 6})
	return nums[len(nums)/2]
}
