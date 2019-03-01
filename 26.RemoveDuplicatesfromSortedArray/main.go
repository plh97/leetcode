package removeduplicates

func removeDuplicates(nums []int) int {
	constainer := make(map[int]int)

	for i := range nums {
		constainer[nums[i]] = nums[i]
	}
	return len(constainer)
}
