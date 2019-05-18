package addStrings

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	res := ""

	addPoint := false

	for i := len(num1) - 1; i >= 0; i-- {

		currentPoint := 0

		if addPoint {
			currentPoint++
			addPoint = false
		}
		aVal, _ := strconv.Atoi(num1[i : i+1])
		bVal := 0
		bValIndex := len(num2) - len(num1) + i
		if bValIndex > -1 {
			bVal, _ = strconv.Atoi(num2[bValIndex : bValIndex+1])
		}
		currentPoint += aVal + bVal
		if currentPoint > 9 {
			currentPoint -= 10
			addPoint = true
		}
		res = strconv.Itoa(currentPoint) + res
	}

	if addPoint {
		res = "1" + res
	}

	return res
}
