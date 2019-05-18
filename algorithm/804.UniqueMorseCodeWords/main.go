package uniqueMorseRepresentations

func uniqueMorseRepresentations(words []string) int {
	mapp := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
		"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--..",
	}

	res := []string{}
	for i := range words {
		// words[i] =
		tempWord := ""
		for j := range words[i] {
			index := words[i][j] - 97
			tempWord += mapp[index]
		}
		if indexOf(res, tempWord) == -1 {
			res = append(res, tempWord)
		}
	}
	return len(res)
}

func indexOf(nums []string, val string) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
