package singleNumber

func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := range nums {
		a = (a ^ nums[i]) &^ b
		b = (b ^ nums[i]) &^ a
	}
	return a
}
