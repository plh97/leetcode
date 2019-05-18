package nextGreatestLetter

func nextGreatestLetter(letters []byte, target byte) byte {
	for i := range letters {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}
