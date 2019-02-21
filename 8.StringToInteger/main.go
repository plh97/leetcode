package myAtoi

import (
	"math"
	"strconv"
	"strings"
)

func clean(s string) string {
	for x := 0; x < len(s); x++ {
		switch s[x : x+1] {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			continue
		case "-":
			if x == 0 {
				continue
			} else {
				return s[:x]
			}
		case "+":
			if x == 0 {
				continue
			} else {
				return s[:x]
			}
		case ".":
			return s[:x]
		default:
			return s[:x]
		}
	}
	return s
}

func myAtoi(str string) int {
	str = clean(strings.TrimSpace(str))
	tempStr, _ := strconv.Atoi(str)
	if tempStr > math.MaxInt32 {
		return math.MaxInt32
	}
	if tempStr < math.MinInt32 {
		return math.MinInt32
	}
	return tempStr
}
