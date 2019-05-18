package reverseOnlyLetters

import "unicode"

func reverseOnlyLetters(S string) string {

	l, r := 0, len(S)-1
	for l < r {
		runeA := []rune(S)
		if unicode.IsLetter(runeA[l]) &&
			unicode.IsLetter(runeA[r]) {
			// 都是字符串
			runeA[r], runeA[l] = runeA[l], runeA[r]
			S = string(runeA)
			l++
			r--
		} else if unicode.IsLetter(runeA[l]) {
			r--
		} else if unicode.IsLetter(runeA[r]) {
			l++
		}else{
			l++
			r--
		}
	}
	return S
}
