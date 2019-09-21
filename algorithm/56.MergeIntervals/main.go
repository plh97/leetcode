package merge

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// sortArray(intervals, 0)
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			a := intervals[i]
			b := intervals[j]
			if a[1] >= b[0] && a[0] <= b[1] {
				// 有重叠
				arr := []int{
					a[0],
					a[1],
					b[0],
					b[1],
				}
				sort.Ints(arr)

				intervals[i] = []int{arr[0], arr[3]}
				// delete j
				intervals = append(intervals[:j], intervals[j+1:]...)
				j = i
			}
		}
	}
	return intervals
}

// pop sort
// func sortArray(arr [][]int, index int) [][]int {
// 	if index > len(arr) {
// 		return arr
// 	}
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		ei := arr[i]
// 		for j := i + 1; j < len(arr); j++ {
// 			ej := arr[j]
// 			if ei[index] > ej[index] {
// 				arr[i], arr[j] = arr[j], arr[i]
// 			}
// 		}
// 	}
// 	return arr
// }
