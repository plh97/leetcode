package foursum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				for m := k + 1; m < l; m++ {
					currentArray := []int{nums[i], nums[j], nums[k], nums[m]}
					// sort.Ints(currentArray)
					if nums[i]+nums[j]+nums[k]+nums[m] == target && isRepeat(res, currentArray) {
						res = append(res, currentArray)
					}
				}
			}
		}
	}
	return res
}

func isRepeat(arrs [][]int, arr []int) bool {
	for i := range arrs {
		isSame := true
		for j := 0; j < len(arr); j++ {
			if arr[j] != arrs[i][j] {
				isSame = false
			}
		}
		if isSame {
			return false
		}
	}
	return true
}
