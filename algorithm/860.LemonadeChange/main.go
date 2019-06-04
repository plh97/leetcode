package lemonadeChange

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _, e := range bills {
		switch e {
		case 5:
			five++
		case 10:
			if five > 0 {
				// 付得起
				five--
				ten++
			} else {
				return false
			}
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
