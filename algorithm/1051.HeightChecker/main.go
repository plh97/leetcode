package heightChecker

import (
	"sort"
)

func heightChecker(heights []int) int {
	res := 0

	memo := make([]int, len(heights))
	copy(memo, heights)
	sort.Sort(sort.IntSlice(memo))
	for i, e := range memo {
		if e != heights[i] {
			res++
		}
	}
	return res
}
