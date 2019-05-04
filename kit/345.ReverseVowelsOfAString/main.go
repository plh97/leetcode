package reverseVowels

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	runes := []rune(s)
	for l < r {
		if isVowels(runes[l]) && isVowels(runes[r]) {
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		} else if isVowels(runes[l]) {
			r--
		} else if isVowels(runes[r]) {
			l++
		} else {
			l++
			r--
		}
	}
	return string(runes)
}

func isVowels(s rune) bool {
	if s == 'a' ||
		s == 'e' ||
		s == 'i' ||
		s == 'u' ||
		s == 'o' ||
		s == 'A' ||
		s == 'E' ||
		s == 'I' ||
		s == 'U' ||
		s == 'O' {
		return true
	}
	return false
}
