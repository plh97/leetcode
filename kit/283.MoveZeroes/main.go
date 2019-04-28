package moveZeroes

import "fmt"

func moveZeroes(nums []int) []int {
	i, j := 0, 1
	for j < len(nums) {
		fmt.Println(nums)
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[i] != 0 && nums[j] == 0 {
			i++
		} else if nums[i] == 0 && nums[j] == 0 {
			j++
		} else if nums[i] != 0 && nums[j] != 0 {
			j += 2
			i += 2
		}
	}
	return nums
}
