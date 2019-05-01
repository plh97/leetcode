package matrixReshape

func matrixReshape(nums [][]int, r int, c int) [][]int {
	res := []int{}
	ress := [][]int{}
	index := 0
	for i := range nums {
		for j := range nums[i] {
			res = append(res, nums[i][j])
		}
	}
	if r*c != len(res) {
		return nums
	}
	//index

	for i := 0; i < r; i++ {
		ress = append(ress, []int{})
		for j := 0; j < c; j++ {
			ress[i] = append(ress[i], res[index])
			index++
		}
	}

	return ress
}
