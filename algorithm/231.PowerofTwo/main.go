package isPowerOfTwo

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	for n != 0 {
		if n == 2 {
			return true
		}
		if n%2 == 0 {
			n /= 2
		} else {
			return false
		}
	}
	return false
}
