package topKFrequent

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	Map := make(map[int]int, k)
	for _, e := range nums {
		Map[e]++
	}
	for ; k > 0; k-- {
		ress := 0
		index := -1
		for i, e := range Map {
			if index < e {
				index = e
				ress = i
			}
		}
		delete(Map, ress)
		res = append(res, ress)
	}
	return res
}
