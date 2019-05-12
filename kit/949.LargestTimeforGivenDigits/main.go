package largestTimeFromDigits

import (
	"strconv"
)

func largestTimeFromDigits(A []int) string {
	res := [2]int{-23, -23}
	random := getRandom(4)
	for i := range random {
		// hour
		a, b := A[random[i][0]-1], A[random[i][1]-1]
		hour := 10*a + b
		a, b = A[random[i][2]-1], A[random[i][3]-1]
		min := 10*a + b
		if hour < 0 || min < 0 || hour > 23 || min > 59 {
			continue
		}
		if hour > res[0] ||
			(hour == res[0] && min > res[1]) {
			res[0] = hour
			res[1] = min
		}
	}
	ress := ""
	if res[0] < 10 && res[0] >= 0 {
		ress += "0" + strconv.Itoa(res[0]) + ":"
	} else if res[0] >= 10 {
		ress += strconv.Itoa(res[0]) + ":"
	}
	if res[1] < 10 && res[1] >= 0 {
		ress += "0" + strconv.Itoa(res[1])
	} else if res[1] >= 10 {
		ress += strconv.Itoa(res[1])
	}
	return ress
}

var dp map[int][][]int = make(map[int][][]int)

func getRandom(n int) [][]int {
	if res, ok := dp[n]; ok {
		return res
	}
	if n == 2 {
		return [][]int{
			[]int{1, 2},
			[]int{2, 1},
		}
	}
	res := [][]int{}
	for i := 1; i <= n; i++ {
		_lowerRes := getRandom(n - 1)
		lowerRes := [][]int{}
		for i := range _lowerRes {
			var temp = make([]int, len(_lowerRes[i]))
			copy(temp, _lowerRes[i])
			lowerRes = append(lowerRes, temp)
		}
		for j := range lowerRes {
			for k := range lowerRes[j] {
				if lowerRes[j][k] == i {
					lowerRes[j][k] = n
				}
			}
			lowerRes[j] = append(lowerRes[j], i)
		}
		res = append(res, lowerRes...)
	}
	dp[n] = res
	return res
}
