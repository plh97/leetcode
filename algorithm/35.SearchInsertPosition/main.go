package searchInsert

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
