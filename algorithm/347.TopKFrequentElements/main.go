package numberOfArithmeticSlices

func numberOfArithmeticSlices(A []int) int {
	res := 0
	long := 0
	for i := 1; i < len(A)-1; i++ {
		// i-1 i i+1
		if A[i]-A[i-1] == A[i+1]-A[i] {
			long++
			res += long
		} else {
			long = 0
		}
	}
	return res
}
