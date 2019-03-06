package lettercombinations

var mapList map[string][]string = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits)==0 {
		return []string{}
	}
	res := []string{""}
	for i := range digits {
		strings := mapList[digits[i:i+1]]
		res = multipartArray(res, strings)
	}
	return res
}

func multipartArray(res []string, b []string) []string {
	result := []string{}
	for i := range res {
		for j := range b {
			result = append(result, res[i]+b[j])
		}
	}
	return result
}
