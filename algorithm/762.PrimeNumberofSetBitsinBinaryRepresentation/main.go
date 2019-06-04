package lastStoneWeight

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		len := len(stones)
		dx := stones[len-1] - stones[len-2]
		if dx == 0 {
			stones = stones[:len-2]
		} else {
			stones = append(stones[:len-2], dx)
		}
		sort.Ints(stones)
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
