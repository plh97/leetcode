package numTilePossibilities

func numTilePossibilities(tiles string) int {
	res := []string{""}

	for _, tile := range tiles {
		for _, ee := range res {
			for i := 0; i <= len(ee); i++ {
				temp := ee[:i] + string(tile) + ee[i:]
				res = append(res, temp)
			}
		}
	}
	Map := make(map[string]bool, len(res))
	for _, e := range res {
		Map[string(e)] = true
	}
	return len(Map) - 1
}

// func subsets(nums []int) [][]int {
// 	res := [][]int{[]int{}}
// 	for _, e := range nums {
// 		for _, ee := range res {
// 			res = append(res, append([]int{e}, ee...))
// 		}
// 	}
// 	return res
// }
