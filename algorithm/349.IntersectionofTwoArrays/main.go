package intersection

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] && indexOf(res, nums1[i]) == -1 {
				res = append(res, nums1[i])
			}
		}
	}
	return res
}

func indexOf(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
