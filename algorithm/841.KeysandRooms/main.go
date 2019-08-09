package canVisitAllRooms

func canVisitAllRooms(rooms [][]int) bool {
	pool := []int{0}
	hasSeen := make(map[int]int)
	hasSeen[0] = 0
	for len(pool) > 0 {
		nextStep := rooms[pool[0]]
		for _, e := range nextStep {
			// 下一步如果未见过
			if _, ok := hasSeen[e]; !ok {
				hasSeen[e] = e
				pool = append(pool, e)
			}
		}
		pool = pool[1:]
	}
	return len(hasSeen) == len(rooms)
}
