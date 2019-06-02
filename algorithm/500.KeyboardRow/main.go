package findWords

import "strings"

func findWords(words []string) []string {
	mp := map[byte]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}
	for ii := 0; ii < len(words); ii++ {
		word := words[ii]
		var start int
	JLoop:
		for i, letter := range word {
			if i == 0 {
				start = mp[strings.ToLower(string(letter))[0]]
			}
			if line, ok := mp[strings.ToLower(string(letter))[0]]; ok {
				if line != start {
					words = append(words[:ii], words[ii+1:]...)
					ii--
					break JLoop
				}
			}
		}
	}
	return words
}
