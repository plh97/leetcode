package searchRange

func searchRange(nums []int, target int) []int {
	firstIndex := indexOf(nums, target)
	lastIndex := lastIndexOf(nums, target)
	return []int{firstIndex, lastIndex}
}

func indexOf(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}

func lastIndexOf(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
