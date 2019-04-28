package findDisappearedNumbers

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for i:=range nums {
		if indexOf(nums,nums[i]) != i {
			res = append(res,i)
		}
	}
	return res
}

func indexOf(nums []int, val int) int {
	for i:=range nums{
		if nums[i] == val {
			return i
		}
	}
	return -1
}