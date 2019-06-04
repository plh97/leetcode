package largestSumAfterKNegations

func largestSumAfterKNegations(A []int, K int) int {
	for K > 0 {
		i := findMin(A)
		A[i] *= -1
		K--
	}
	res := 0
	for _, e := range A {
		res += e
	}
	return res
}

func findMin(A []int) int {
	temp := 0
	for i, e := range A {
		if A[temp] > e {
			temp = i
		}
	}
	return temp
}
