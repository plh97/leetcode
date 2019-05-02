package flipAndInvertImage

func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		// flip
		j, k := 0, len(A[i])-1
		for j <= k {
			if j == k {
				A[i][j] = A[i][j] ^ 1
			} else {
				A[i][j], A[i][k] = A[i][k]^1, A[i][j]^1
			}
			j++
			k--
		}
	}
	return A
}
