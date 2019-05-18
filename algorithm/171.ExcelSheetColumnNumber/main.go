package titleToNumber

func titleToNumber(s string) int {
	res := 0
	for i := range s {
		if i > 0 {
			res *= 26
		}
		res += int(s[i]-64)
	}
	return res
}
