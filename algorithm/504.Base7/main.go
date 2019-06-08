package convertToBase7

import "strconv"

func convertToBase7(num int) string {
	res := ""

	if num < 0 {
		num *= -1
		for num > 0 {
			res = strconv.Itoa(num%7) + res
			num /= 7
		}
		res = "-" + res
	} else if num > 0 {
		for num > 0 {
			res = strconv.Itoa(num%7) + res
			num /= 7
		}
	} else {
		return "0"
	}
	return res
}
