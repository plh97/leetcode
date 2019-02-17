package FindMedianSortedArrays

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	for i := 0; i < len(nums3); i++ {
		for j := i; j < len(nums3)-1; j++ {
			if nums3[j] > nums3[j+1] {
				nums3[j], nums3[j+1] = nums3[j+1], nums3[j]
			}
		}
	}
	fmt.Println(nums3, (len(nums3)-1)/2)

	if len(nums3)%2 == 0 {
		return float64(nums3[(len(nums3)-1)/2]+nums3[(len(nums3))/2])/2
	} else {
		return float64(nums3[(len(nums3)-1)/2])
	}
}

// func main() {
// 	fmt.Println(findMedianSortedArrays("aab"))
// }
