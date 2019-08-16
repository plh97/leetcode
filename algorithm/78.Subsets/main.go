package subsets

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	for _, e := range nums {
		currentOutputLength := len(res)
		for i := 0; i < currentOutputLength; i++ {
			tempSlice := make([]int, len(res[i]))
			copy(tempSlice, res[i])
			tempSlice = append(tempSlice, e)
			res = append(res, tempSlice)
		}
	}

	return res
}
