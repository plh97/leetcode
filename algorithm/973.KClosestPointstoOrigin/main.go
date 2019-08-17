package kClosest

import (
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	ress := []int{}
	res := [][]int{}
	for _, e := range points {
		dis := e[0]*e[0] + e[1]*e[1]
		ress = append(ress, dis)
	}
	sort.Ints(ress)
	ress = ress[:K]
	for _, e1 := range ress {
		for j := 0; j < len(points); j++ {
			point := points[j]
			if e1 == point[0]*point[0]+point[1]*point[1] {
				points = append(points[:j], points[j+1:]...)
				res = append(res, point)
			}
		}
	}
	return res
}
