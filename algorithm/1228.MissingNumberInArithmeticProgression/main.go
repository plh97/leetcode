package missingnumber

import "math"

func missingNumber(arr []int) int {
	if len(arr) < 3 {
		return -1
	}
	intMap := make(map[int]int)
	for i := 1; i < len(arr); i++ {
		dx := arr[i] - arr[i-1]
		if _, ok := intMap[dx]; ok {
			intMap[dx]++
		} else {
			intMap[dx] = 1
		}
	}
	max := 0
	arr1 := []int{}
	for i := range intMap {
		arr1 = append(arr1, i)
	}
	if intMap[arr1[0]] > intMap[arr1[1]] {
		max = arr1[0]
	} else if intMap[arr1[0]] < intMap[arr1[1]] {
		max = arr1[1]
	} else {
		if math.Abs(float64(arr1[0])) < math.Abs(float64(arr1[1])) {
			max = arr1[0]
		} else {
			max = arr1[1]
		}
	}
	res := 0
	for i := 1; i < len(arr); i++ {
		if max != (arr[i] - arr[i-1]) {
			res = arr[i-1] + max
		}
	}

	return res
}
