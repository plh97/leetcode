package countSegments

func countSegments(s string) int {
	res := 0
	start := false
	for i := range s {
		currentString := s[i : i+1]
		if currentString != " " {
			start = true
		} else {
			if start {
				res++
			}
			start = false
		}
		if currentString != " " && i == len(s)-1 {
			if start {
				res++
			}
		}
	}
	return res
}
