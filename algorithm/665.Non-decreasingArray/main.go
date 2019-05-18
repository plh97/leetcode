package checkPossibility

func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		numArr := copyArr(nums, i)
		// 随机删除任意一个数字
		if isInc(numArr) {
			return true
		}
	}
	return false
}

func isInc(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func copyArr(nums []int, ii int) []int {
	newArr := []int{}
	for i := range nums {
		if i == ii {
			continue
		}
		newArr = append(newArr, nums[i])
	}
	return newArr
}
