package canConstruct

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for i := range ransomNote {
		currentString := ransomNote[i : i+1]
		currentIndex := strings.Index(magazine, currentString)
		if currentIndex > -1 {
			magazine = magazine[:currentIndex] + magazine[currentIndex+1:]
		} else {
			return false
		}
	}
	return true
}
