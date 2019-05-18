package findUnsortedSubarray

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	tempNums := copyArr(nums)
	start := 0
	end := 0
	sort.Ints(tempNums)
	for i := range nums {
		if tempNums[i] != nums[i] {
			start = i
			break
		}
	}
	for i := len(nums) - 1; i > 0; i-- {
		if tempNums[i] != nums[i] {
			end = i
			break
		}
	}
	if end > start{
		return end - start + 1
	}
	return 0
}
func copyArr(nums []int) []int {
	newArr := []int{}
	for i := range nums {
		newArr = append(newArr, nums[i])
	}
	return newArr
}
