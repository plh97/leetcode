package sortArrayByParity

func sortArrayByParity(A []int) []int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i]%2 == 1 && A[j]%2 == 0 {
				// i 偶数
				// j 偶数
				A[i], A[j] = A[j], A[i]
			} else if A[i]%2 == 1 && A[j]%2 == 1 {
				// i 偶数
				// j 奇数
				continue
			} else {
				break
			}
		}
	}

	return A
}
