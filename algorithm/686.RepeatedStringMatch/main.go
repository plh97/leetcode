package repeatedStringMatch

import (
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	res := ""
	for i := 0; i < len(B)/len(A)+3; i++ {
		if strings.Index(res, B) > -1 {
			// 找到了
			return i
		} else if i == len(B)/len(A)+2 && strings.Index(res, B) == -1 {
			// 最后一次都找不到
			res = ""
		} else {
			// 未找到
			res += A
		}
	}
	if len(res) == 0 {
		return -1
	}
	return len(res) / len(A)
}
