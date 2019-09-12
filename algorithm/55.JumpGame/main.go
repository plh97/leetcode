package canJump

func canJump(nums []int) bool {
	res := helper(nums)
	return res < 3
}

func helper(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	stepRange := nums[0]
	if stepRange >= len(nums)-1 {
		return 0
	} else {
		container := []int{}
		for i := 1; i <= stepRange; i++ {
			container = append(container, 1+helper(nums[i:]))
		}
		return min(container)
	}
}

func min(nums []int) int {
	res := nums[0]
	for _, e := range nums {
		if res > e {
			res = e
		}
	}
	return res
}
