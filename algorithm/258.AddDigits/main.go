package addDigits

func addDigits(num int) int {
	for num > 9 {
		init := 0
		for num > 0 {
			init += num % 10
			num /= 10
		}
		num = init
	}
	return num
}
