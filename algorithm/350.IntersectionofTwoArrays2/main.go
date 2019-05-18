package intersect

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	n, m := len(nums1), len(nums2)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				nums1 = append(nums1[:i], nums1[i+1:]...)
				nums2 = append(nums2[:j], nums2[j+1:]...)
				n--
				m--
				i--
				j--
				break
			}
		}
	}
	return res
}
