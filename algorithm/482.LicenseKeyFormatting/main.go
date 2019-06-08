package licenseKeyFormatting

import (
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	S = strings.Replace(S, "-", "", len(S))
	S = strings.ToUpper(S)
	if K == 1 {
		for i := len(S) - 1; i >= 0; i-- {
			if i > 0 && (len(S)-i)%K == K-1 {
				S = S[:i] + "-" + S[i:]
			}
		}
	} else {
		for i := len(S) - 1; i >= 0; i-- {
			if i > 0 && (len(S)-i+1)%(K+1) == 0 {
				S = S[:i] + "-" + S[i:]
			}
		}
	}
	return S
}

// func newLen(S string) int {
// 	return len(strings.Replace(S, "-", "", len(S)))
// }
