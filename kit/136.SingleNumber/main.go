package singleNumber

func singleNumber(nums []int) int {
	for i := range nums {
		l_index:=indexOf(nums, nums[i])
		r_index:=lastIndexOf(nums, nums[i])

		if l_index == r_index {
			return nums[i]
		}
	}
	return 1
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
