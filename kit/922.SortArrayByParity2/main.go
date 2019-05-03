package sortArrayByParityII

func sortArrayByParityII(A []int) []int {
	for l := 0; l < len(A)-1; l++ {
		for r := l+1; r < len(A); r++ {
			if A[l]%2 != l%2 && A[r]%2 != r%2 && A[l]%2 != A[r]%2 {
				A[l], A[r] = A[r], A[l]
				l++
				break
			} else if A[l]%2 == l%2 {
				break
			}
		}
	}
	return A
}
