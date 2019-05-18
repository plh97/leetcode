package repeat

func repeat(a string) map[string]int {
	var res map[string]int = map[string]int{}
	for i := range a {
		res[a[i:i+1]]++
	}
	return res
}