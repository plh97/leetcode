package multiple

// LCM -> 公约
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCD -> 公倍
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func multiple(a, b int) [2]int {
	return [2]int{gcd(a, b), lcm(a, b)}
}

// func main() {
// 	a, _ := strconv.ParseInt(os.Args[1], 0, 64)
// 	b, _ := strconv.ParseInt(os.Args[2], 0, 64)
// 	_a, _b := int(a), int(b)

// 	fmt.Println(multiple(_a, _b))
// }
