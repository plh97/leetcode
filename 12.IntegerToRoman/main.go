package intToRoman

var RomanMap [][]string = [][]string{
	[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	[]string{"", "M", "MM", "MMM", "CD", "D", "DC", "DCC", "DCCC", "DCCCC"},
}

func intToRoman(num int) string {
	return RomanMap[3][num%10000/1000] +
		RomanMap[2][num%1000/100] +
		RomanMap[1][num%100/10] +
		RomanMap[0][num%10/1]
}
