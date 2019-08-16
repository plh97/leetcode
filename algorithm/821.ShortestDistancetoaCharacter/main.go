package shortestToChar

func shortestToChar(S string, C byte) []int {
	res := []int{}
	for i := range S {
		if S[i] == C {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	}
	count := 9999
	for i, e := range res {
		// 顺着
		if e == 0 {
			count = 0
		} else {
			count++
			res[i] = count
		}
	}
	for i := len(res)-1; i >= 0; i-- {
		// 逆序着
		if res[i] == 0 {
			count = 0
		} else {
			count++
			if res[i] > count {
				res[i] = count
			}
		}
	}
	return res
}
