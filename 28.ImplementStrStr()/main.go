package strStr

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	} else if haystack == needle {
		return 0
	} else {
		for i := 0; i < len(haystack)-len(needle); i++ {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
