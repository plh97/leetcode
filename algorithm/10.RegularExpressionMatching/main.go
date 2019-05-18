package isMatch

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	dp := make([][]bool, sLen+1)
	// 获取完整的动态规划表
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	// 将第一行中*变成1
	for i := 1; i < pLen && dp[0][i-1]; i += 2 {
		if p[i] == '*' {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < sLen; i++ {
		for j := 0; j < pLen; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				// 此处需要判断 \w*, *前面的单词是否能与s[i]匹配,
				if p[j-1] == s[i] || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i][j+1] ||
						dp[i+1][j] ||
						dp[i+1][j-1]
				} else {
					// 否则p[i-1:i+1] \w*只能被当做 "" 来消耗
					// 因此,需要去查看上上个dp是否为true
					dp[i+1][j+1] = dp[i+1][j-1]
				}
			}
		}
	}
	return dp[sLen][pLen]

}
