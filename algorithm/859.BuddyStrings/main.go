package buddyStrings

func buddyStrings(A string, B string) bool {
	if A == B {
		for i := range A {
			if i > 0 && A[i] == A[0] {
				return true
			}
		}
		return false
	} else {
		l, r := 0, len(A)-1
		for l < r {
			if A[r] != B[r] && A[l] != B[l] {
				runeA := []rune(A)
				runeA[r], runeA[l] = runeA[l], runeA[r]
				A = string(runeA)
				break
			} else {
				if A[l] == B[l] {
					l++
				}
				if A[r] == B[r] {
					r--
				}
			}
		}
	}
	return A == B
}
