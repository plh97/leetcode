package findDisappearedNumbers

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for i := 1; i <= len(nums); i++ {
		res = append(res, i)
	}
	for i := range nums {
		index := indexOf(res, nums[i])
		if index > -1 {
			res = append(res[:index], res[index+1:]...)
		}
	}
	return res
}

func indexOf(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
