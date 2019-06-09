package findErrorNums

func findErrorNums(nums []int) []int {
	Map := make(map[int]bool, len(nums))
	res := []int{0, 0}
	sum := 0
	for _, e := range nums {
		sum += e
		if ok := Map[e]; ok {
			res[0] = e
		}
		Map[e] = true
	}
	len := len(nums)
	res[1] = len*(len+1)/2 - (sum - res[0])
	return res
}
