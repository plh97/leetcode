package findTheDifference

func findTheDifference(s string, t string) byte {
	var c byte = t[len(t)-1]
	for i := range s {
		c ^= t[i]
		c ^= s[i]
	}
	return c
}

// func findTheDifference(s string, t string) byte {
// 	var c byte
// 	for i := range s {
// 		c ^= s[i]
// 	}
// 	for i := range t {
// 		c ^= t[i]
// 	}
// 	return c
// }

// func findTheDifference(s string, t string) byte {
// 	for i := 0; i < len(t); i++ {
// 		index := strings.Index(s, t[i:i+1])
// 		if index == -1 {
// 			return t[i]
// 		} else {
// 			s = s[:index] + s[index+1:]
// 		}
// 	}
// 	return 'a'
// }
