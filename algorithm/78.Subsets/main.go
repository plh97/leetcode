package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 1, 1024)
	for _, e := range nums {
		for _, ee := range res {
			res = append(res, append([]int{e}, ee...))
		}
	}
	return res
}
