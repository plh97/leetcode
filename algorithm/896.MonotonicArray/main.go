package isMonotonic

func isMonotonic(A []int) bool {
	inc := true
	dec := true
	// inc
	for i := 1; i < len(A); i++ {
		if A[i-1] > A[i] {
			inc = false
		} else if A[i-1] < A[i] {
			dec = false
		}
		if inc == false && dec == false {
			return false
		}
	}
	return inc || dec
}
