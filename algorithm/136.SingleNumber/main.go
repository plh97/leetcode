package singleNumber

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res ^= nums[i]
	}
	return res
}
