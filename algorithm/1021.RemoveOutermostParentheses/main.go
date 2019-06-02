package removeOuterParentheses

import "bytes"

func removeOuterParentheses(S string) string {
	sb := bytes.Buffer{}
	count := 0
	for l, r := 0, 0; r < len(S); r++ {
		e := S[r]
		if e == '(' {
			count++
		} else if e == ')' {
			count--
		}
		if count == 0 {
			sb.WriteString(S[l+1 : r])
			l = r + 1
		}
	}
	return sb.String()
}
