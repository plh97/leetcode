package getRow
 
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := [][]int{
		[]int{1},
	}
	for i := 1; i <= rowIndex; i++ {
		item := []int{1}
		for j := 1; j < i; j++ {
			item = append(item, res[i-1][j-1]+res[i-1][j])
		}
		item = append(item, 1)
		res = append(res, item)
	}
	return res[rowIndex]
}
