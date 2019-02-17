package longestPalindrome

func isPalindromicSubstring(s string) string {
	if len(s) == 1 {
		return s
	}
	if len(s)%2 == 0 {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return ""
			}
		}
	} else {
		// 奇数
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return ""
			}
		}
	}
	return s
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var longString string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			// 检测这个字符串是否是回文
			temp := isPalindromicSubstring(s[i:j])
			if len(longString) < len(temp) {
				longString = temp
			}
		}
	}
	return longString
}
