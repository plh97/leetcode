package nextPermutation

func nextPermutation(nums []int) {
	var (
		l = len(nums) - 1
		r = len(nums) - 1
	)
	for l-1 >= 0 && nums[l] <= nums[l-1] {
		l--
	}
	if l == 0 {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
		return
	}
	k := r
	for nums[k] <= nums[l-1] {
		k--
	}
	nums[k], nums[l-1] = nums[l-1], nums[k]
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return
}
