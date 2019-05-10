package mySqrt

func mySqrt(x int) int {
	xx := 1
	for xx*xx <= x {
		xx++
	}
	return xx-1
}
