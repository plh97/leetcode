package addBinary


func addBinary(a string, b string) string {
	res := ""
	if len(a) < len(b) {
		a, b = b, a
	}
	addPlus := false
	for i := len(a) - 1; i >= 0; i-- {

		aVal := a[i : i+1]

		bVal := "0"

		if len(b)-len(a)+i >= 0 {
			bVal = b[len(b)-len(a)+i : len(b)-len(a)+i+1]
		} else {
			bVal = "0"
		}
		currentPoint := 0

		if addPlus {
			currentPoint++
			addPlus = false
		}
		if aVal == "0" && bVal == "1" || aVal == "1" && bVal == "0" {
			currentPoint++
		} else if aVal == "1" && bVal == "1" {
			currentPoint += 2
		}

		if currentPoint > 1 {
			currentPoint -= 2
			addPlus = true
		}

		if currentPoint == 0 {
			res = "0" + res
		} else {
			res = "1" + res
		}

	}

	if addPlus {
		res = "1" + res
	}

	return res
}
