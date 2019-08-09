package dailyTemperatures

func dailyTemperatures(T []int) []int {
	for i := range T {
		if i == len(T)-1 {
			T[i] = 0
		}
	JLoop:
		for j := i + 1; j < len(T); j++ {
			if T[i] < T[j] {
				T[i] = j - i
				break JLoop
			} else if j == len(T)-1 && T[i] >= T[j] {
				T[i] = 0
				break JLoop
			}
		}
	}
	return T
}
