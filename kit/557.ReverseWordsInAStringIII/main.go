package reverseWords

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for i := range arr {
		arr[i] = swep(arr[i])
	}
	res := ""
	for i := range arr {
		res += arr[i]
		if i != len(arr)-1 {
			res += " "
		}
	}
	return res
}

func swep(ss string) string {
	l, r := 0, len(ss)-1
	s := []rune(ss)
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return string(s)
}
