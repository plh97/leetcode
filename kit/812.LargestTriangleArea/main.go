package largestTriangleArea

import (
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	res := 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				x, y, z := points[i], points[j], points[k]
				x1 := x[0]
				y1 := x[1]
				x2 := y[0]
				y2 := y[1]
				x3 := z[0]
				y3 := z[1]
				newArea := math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2) / 2.0)
				res = math.Max(res, newArea)
			}
		}
	}
	return res
}
