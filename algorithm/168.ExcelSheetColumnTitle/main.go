package convertToTitle

func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		v := n % 26 // 余数
		n /= 26     // 向前进一位
		if v == 0 {
			v = 26
			n--
		}
		res = string(v+64) + res
	}
	return res
}

// 701

// 701

// A*26^2 + B*26^1 + C*26^0 = 701
// ABC = 701

// 701 / 26
