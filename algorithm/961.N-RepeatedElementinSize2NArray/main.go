package repeatedNTimes

func repeatedNTimes(A []int) int {
	hasSeen := [10000]bool{}
	res := 0
	for _, res := range A {
		if hasSeen[res] {
			return res
		}
		hasSeen[res] = true
	}
	return res
}
