package findRadius

import "math"

// 供暖点  ->  暖源
func findRadius(houses []int, heaters []int) int {
	res := 0
	for i := 0; i < len(houses); i++ {
		dis := math.MaxInt64
		for j := 0; j < len(heaters); j++ {
			dis = Min(dis, Abs(houses[i]-heaters[j]))
		}
		res = Max(res, dis)
	}
	return res
}

func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
