package isValid

var Map map[string]string = map[string]string{
	"[": "]",
	"(": ")",
	"{": "}",
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

func isContain(stack []string, s string) bool {
	for i := range stack {
		if _, ok := Map[stack[i]]; ok && s == Map[stack[i]] {
			return true
		}
	}
	return false
}
