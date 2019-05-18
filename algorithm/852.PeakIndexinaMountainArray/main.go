package peakIndexInMountainArray

func peakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] &&
			A[i] > A[i+1] {
			return i
		}
	}
	return -1
}
