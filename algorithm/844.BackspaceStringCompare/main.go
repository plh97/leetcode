package backspaceCompare

func backspaceCompare(S string, T string) bool {
	return dec(S) == dec(T)
}

func dec(T string) string {
	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			if i > 0 {
				T = T[:i-1] + T[i+1:]
				i -= 2
			} else {
				T = T[:i] + T[i+1:]
				i--
			}
		}
	}
	return T
}
