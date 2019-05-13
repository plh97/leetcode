package bitwiseComplement

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	res := 0
	for i := 1; N > 0; i *= 2 {
		if N%2 == 0 {
			res += i
		}
		N /= 2
	}
	return res
}
