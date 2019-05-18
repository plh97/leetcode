package reverseString

func reverseString(s []byte) []byte {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}
