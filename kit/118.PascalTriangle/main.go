package generate

func generate(numRows int) [][]int {
	if numRows==0 {
		return [][]int{}
	}
	res := [][]int{
		[]int{1},
	}
	for i := 1; i < numRows; i++ {
		item := []int{1}
		for j := 1; j < i; j++ {
			item = append(item, res[i-1][j-1]+res[i-1][j])
		}
		item = append(item, 1)
		res = append(res, item)
	}
	return res
}
