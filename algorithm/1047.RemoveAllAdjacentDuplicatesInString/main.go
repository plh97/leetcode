package removeDuplicates

func removeDuplicates(S string) string {
	isDiff := true
	for isDiff {
		isDiff = false
		for i := 0; i < len(S)-1; i++ {
			if S[i] == S[i+1] {
				S = S[:i] + S[i+2:] // 去掉i 和 i+1
				i--
				isDiff = true
			}
		}

	}
	return S
}
