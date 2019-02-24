package main

import "fmt"

func main() {
	fmt.Println(isMatch("aaaaaa", "a*a"))
}

func isMatch(s string, p string) bool {
	sSize := len(s)
	pSize := len(p)
	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}
	dp[0][0] = true
	/*
		a * a a a
		x ^ x x x
	*/
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == s[i] {
				/*
					p[j] 与 s[i] 可以匹配上，
					所以，只要前面匹配，这里就能匹配上
				*/
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				fmt.Println(i, j)
				dp[i+1][j+1] =
					dp[i+1][j-1] ||
						dp[i+1][j] ||
						dp[i][j+1]
			}
		}
	}
	return dp[sSize][pSize]
}
