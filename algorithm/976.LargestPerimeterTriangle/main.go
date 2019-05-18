package largestPerimeter

import (
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Ints(A)
	res := 0
	for i := len(A) - 1; i > 1; i-- {
		if A[i] < A[i-1]+A[i-2] {
			res := A[i] + A[i-1] + A[i-2]
			return res
		}
	}
	return res
}

// func largestPerimeter(A []int) int {
// 	res := 0
// 	for i := 0; i < len(A)-2; i++ {
// 		for j := i + 1; j < len(A)-1; j++ {
// 			for k := j + 1; k < len(A); k++ {
// 				arr := []int{
// 					A[i],
// 					A[j],
// 					A[k],
// 				}
// 				sort.Ints(arr)
// 				if arr[1]+arr[0] > arr[2] {
// 					res = max(res, arr[0]+arr[1]+arr[2])
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
