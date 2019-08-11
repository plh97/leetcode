package singleNumber

func singleNumber(nums []int) []int {
	Map := make(map[int]int, len(nums))
	res := []int{}
	for _, e := range nums {
		Map[e]++
	}
	for i, e := range Map {
		if e == 1 {
			res = append(res, i)
		}
	}
	return res
}
