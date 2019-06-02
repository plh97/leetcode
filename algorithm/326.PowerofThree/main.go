package isPowerOfThree

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for n != 0 {
		if n == 3 {
			return true
		}
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
	return false
}
