package sortedSquares

import "sort"

func sortedSquares(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}
