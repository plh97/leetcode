package detectCapitalUse

import "strings"

func detectCapitalUse(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}
