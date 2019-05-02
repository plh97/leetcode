package pivotIndex

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if totalArr(nums[:i]) == totalArr(nums[i+1:]) {
			return i
		}
	}
	return -1
}

func totalArr(nums []int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}
