package singleNumber

func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := range nums {
		a ^= nums[i] &^ b
		b ^= nums[i] &^ a
	}
	return a
}
