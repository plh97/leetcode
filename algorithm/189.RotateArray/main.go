package rotate

func rotate(nums []int, k int) []int {
	for k > 0 {
		k--
		pop(nums)
	}
	return nums
}

func pop(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
}
