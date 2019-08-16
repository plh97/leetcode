package relativeSortArray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := []int{}
	for _, a2 := range arr2 {
		for j, a1 := range arr1 {
			if a2 == a1 {
				res = append(res, a2)
				arr1[j] = -999999
			}
		}
	}
	sort.Ints(arr1)
	for _, e := range arr1 {
		if e != -999999 {
			res = append(res, e)
		}
	}
	return res
}
