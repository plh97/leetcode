package addToArrayForm

func addToArrayForm(A []int, K int) []int {
	B := []int{}
	if K == 0 {
		B = []int{0}
	}
	for K > 0 {
		B = append([]int{K % 10}, B...)
		K = K / 10
	}
	add := 0
	if len(A) < len(B) {
		A, B = B, A
	}
	for i := len(A) - 1; i >= 0; i-- {
		if add == 1 {
			A[i]++
			add = 0
		}
		valB := 0
		if len(A)-len(B) <= i {
			valB = B[i - len(A) + len(B)]
		} else {
			valB = 0
		}
		if valB+A[i]+add > 9 {
			A[i] += valB - 10 + add
			add = 1
		} else {
			A[i] += valB
		}
	}
	if add == 1 {
		return append([]int{1}, A...)
	} else {
		return A
	}

}

