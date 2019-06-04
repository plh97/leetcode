package twoCitySchedCost

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		a, b := costs[i], costs[j]
		return (a[0] - a[1]) < (b[0] - b[1])
	})
	res := 0
	for i, e := range costs {
		if i < len(costs)/2 {
			res += e[0]
		} else {
			res += e[1]
		}
	}
	return res
}
