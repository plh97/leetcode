package countPrimeSetBits

func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		cp := i
		temp := 0
		for cp > 0 {
			if cp%2 == 1 {
				temp++
			}
			cp /= 2
		}
		if isPrime(temp) {
			res++
		}
	}
	return res
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
