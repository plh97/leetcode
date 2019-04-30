package containsNearbyDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		index := indexOf(nums[i+1:], nums[i])+1
		if index > 0 && index <= k {
			return true
		}
	}
	return false
}

func indexOf(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}

// func lastIndexOf(nums []int, val int) int {
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		if nums[i] == val {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func isInner(nums []int, k int) bool {
// 	for i := 1; i < len(nums)-1; i++ {
// 		if nums[i]-nums[0] > k || nums[0]-nums[i] > k {
// 			return false
// 		}
// 	}
// 	return true
// }
