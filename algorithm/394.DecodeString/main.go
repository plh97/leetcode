package decodeString

import "strings"

func decodeString(s string) string {
	stringStack := []string{}
	numberStack := []int{}
	res := ""
	num := 0
	for _, e := range s {
		switch {
		case e >= '0' && e <= '9':
			num = 10*num + int(e-'0')
		case e == '[':
			numberStack = append(numberStack, num)
			num = 0
			stringStack = append(stringStack, res)
			res = ""
		case e == ']':
			strTemp := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			numTemp := numberStack[len(numberStack)-1]
			numberStack = numberStack[:len(numberStack)-1]
			res = strTemp + strings.Repeat(res, numTemp)
		default:
			res += string(e)
		}
	}
	return res
}
