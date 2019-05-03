package validMountainArray

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	n := 2
	// 仅仅当存在一个峰值点, 这个峰值点index!=0 && index != len(A)
	for i := 1; i < len(A)-1; i++ {
		start := A[i] - A[i-1]
		end := A[i+1] - A[i]
		if start > 0 && end < 0 {
			n--
		}
		if start < 0 && end > 0 || start == 0 || end == 0 {
			return false
		}
	}
	return n == 1
}
