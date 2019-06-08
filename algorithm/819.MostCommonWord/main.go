package mostCommonWord

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	Map := make(map[string]int)
	start := 0
	res := ""
	for i := 0; i <= len(paragraph); i++ {
		if i == len(paragraph) || paragraph[i] < 'a' || paragraph[i] > 'z' {
			letter := paragraph[start:i]
			if len(letter) > 0 && indexOf(banned, letter) == -1 {
				if _, ok := Map[letter]; ok {
					Map[letter]++
				} else {
					Map[letter] = 1
				}
				if Map[res] < Map[letter] {
					res = letter
				}
			}
			start = i + 1
		}
	}
	return res
}

func indexOf(nums []string, val string) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
