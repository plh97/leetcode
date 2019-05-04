package canConstruct

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for i := range ransomNote {
		currentString := ransomNote[i : i+1]
		currentIndex := strings.Index(magazine, currentString)
		if currentIndex > -1 {
			// 找到即可去重.
			fmt.Println(currentString)
		} else {
			// 找不到
			return false
		}

	}
	return true
}
