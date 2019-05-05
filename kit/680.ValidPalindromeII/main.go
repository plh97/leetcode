package validPalindrome

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			if isPalindrome(s, l+1, r) {
				return true
			}
			if isPalindrome(s, l, r-1) {
				return true
			}
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
