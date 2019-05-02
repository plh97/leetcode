package transpose

func transpose(nums [][]int) [][]int {
	res := []int{}
	ress := [][]int{}
	index := 0
	c := len(nums)
	r := len(nums[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res = append(res, nums[j][i])
		}
	}
	// if r*c != len(res) {
	// 	return nums
	// }
	for i := 0; i < r; i++ {
		ress = append(ress, []int{})
		for j := 0; j < c; j++ {
			ress[i] = append(ress[i], res[index])
			index++
		}
	}

	return ress
}
