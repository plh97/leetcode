package findLengthOfLCIS

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	for i := range nums {
		res = max(res, fromBeginIncNum(nums[i:]))
	}
	return res
}

func fromBeginIncNum(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return i
		}
	}
	return len(nums)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
