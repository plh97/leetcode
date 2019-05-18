package largeGroupPositions

func largeGroupPositions(S string) [][]int {
	res := [][]int{}
	start := 0
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			if i-start > 2 {
				res = append(res, []int{start, i - 1})
			}
			start = i
		} else if i == len(S)-1 && i-start > 1 {
			res = append(res, []int{start, i})
		}
	}
	return res
}
