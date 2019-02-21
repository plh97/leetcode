package repeat

func repeat(a string) map[string]int {
	var res map[string]int = map[string]int{}
	for i := range a {
		res[a[i:i+1]]++
	}
	return res
}

// func main() {
// 	a, _ := strconv.ParseInt(os.Args[1], 0, 64)
// 	b, _ := strconv.ParseInt(os.Args[2], 0, 64)
// 	_a, _b := int(a), int(b)

// 	fmt.Println(multiple(_a, _b))
// }
