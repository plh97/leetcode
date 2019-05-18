package checkPerfectNumber

func checkPerfectNumber(num int) bool {
	if num == 0 {
		return false
	}
	res := 1
	// 2 对应 num/2
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			res += i + num/i
		}
	}
	return res == num
}
