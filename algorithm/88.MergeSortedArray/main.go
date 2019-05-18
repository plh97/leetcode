package merge

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[:m], nums2[:n]...)

    for i := 0; i < len(nums1)-1; i++ {
        for j := i+1; j < len(nums1); j++ {
            if  nums1[i]>nums1[j]{
                nums1[i],nums1[j] = nums1[j],nums1[i]
            }
        }
    }
	return nums1
}
