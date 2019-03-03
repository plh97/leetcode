package removeduplicates

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
