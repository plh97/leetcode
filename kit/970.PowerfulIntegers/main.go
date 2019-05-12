package powerfulIntegers

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	res := []int{2}
	for i := 3; i <= bound; i++ {
	Exit:
		for j := 0; math.Pow(float64(x), float64(j)) < float64(i) && (x!=1||j==0); j++ {
			for k := 0; math.Pow(float64(y), float64(k)) < float64(i) && (y!=1||k==0); k++ {
				if math.Pow(float64(y), float64(k))+math.Pow(float64(x), float64(j)) == float64(i) {
					res = append(res, i)
					break Exit
				}
			}
		}
	}
	return res
}
