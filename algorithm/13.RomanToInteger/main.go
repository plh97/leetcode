package romanToInt

func romanToInt(s string) int {
	RomanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for len(s) > 0 {

		if len(s) > 1 && RomanMap[s[:1]] < RomanMap[s[1:2]] {
			res += (RomanMap[s[1:2]] - RomanMap[s[:1]])
			s = s[2:] // s 消耗2个字符
		} else {
			res += RomanMap[s[0:1]]
			s = s[1:] // s 消耗1个字符
		}
	}
	return res
}
