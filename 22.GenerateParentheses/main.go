package generateparenthesis

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	n = n * 2
	res := []string{}

	for i := 0; i < (1 << uint(n)); i++ {
		// 列举所有可能
		ss := ""
		for j := n; j > 0; j-- {
			temp := i % (1 << uint(j)) / (1 << uint(j-1))
			if temp > 0 {
				ss += ")"
			} else if temp == 0 {
				ss += "("
			}
		}
		// check it here
		if checkOut(ss) {
			fmt.Println(uint64(i), ss)
			res = append(res, ss)
		}

	}

	return res
}

func checkOut(s string) bool {
	// 不能以)为开头
	if s[0] == ')' {
		return false
	}
	// 不能以(为结尾
	if s[len(s)-1] == '(' {
		return false
	}
	// ()数量对称多
	time := 0
	for i := range s {
		if s[i] == '(' {
			time++
		}
	}
	if time != len(s)/2 {
		return false
	}
	if !isValid(s) {
		return false
	}
	return true
}

func isValid(s string) bool {
	stack := []string{}
	for i := range s {
		ss := s[i : i+1]
		switch s[i] {
		case '[', '(', '{':
			stack = append(stack, ss)
		case ']', ')', '}':
			if len(stack) == 0 {
				return false
			}
			if isContain(stack, ss) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, ss)
			}
		}
	}
	return len(stack) == 0
}

var Map map[string]string = map[string]string{
	"[": "]",
	"(": ")",
	"{": "}",
}

func isContain(stack []string, s string) bool {
	for i := range stack {
		if _, ok := Map[stack[i]]; ok && s == Map[stack[i]] {
			return true
		}
	}
	return false
}
