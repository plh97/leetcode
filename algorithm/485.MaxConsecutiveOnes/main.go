package findMaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	var sum, max int
	for _, v := range nums {
		sum += v
		sum *= v
		if sum > max {
			max = sum
		}
	}
	return max
}

// func findMaxConsecutiveOnes(nums []int) int {
// 	res := 0
// 	count := 1
// 	if len(nums) == 1 && nums[0] == 1 {
// 		return 1
// 	}
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == nums[i+1] {
// 			if nums[i] == 1 {
// 				count++
// 			} else {
// 				continue
// 			}
// 		} else {
// 			count = 1
// 		}
// 		if count > res {
// 			res = count
// 		}
// 	}
// 	return res
// }
