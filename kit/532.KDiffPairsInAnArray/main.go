package findPairs

import "sort"

func findPairs(nums []int, k int) int {
	res := 0
	if k < 0 {
		return 0
	}
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == k {
				res++
				break
			}
		}
	}
	return res
}
