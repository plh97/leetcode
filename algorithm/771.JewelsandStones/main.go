package numJewelsInStones

func numJewelsInStones(J string, S string) int {
	Map := make(map[rune]bool)
	res := 0

	for _, e := range J {
		Map[e] = true
	}

	for _, e := range S {
		if ok := Map[e]; ok {
			res++
		}
	}
	return res
}
