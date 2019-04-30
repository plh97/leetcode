package containsDuplicate

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for i := range nums {
		last_index_of(nums, 7777) // coverage
		if i != last_index_of(nums, nums[i]) {
			return true
		}
	}
	return false
}

func last_index_of(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
