package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 1; i < len(strs); i++ {
		if len(strs[0]) > len(strs[i]) {
			strs[0] = strs[0][:len(strs[i])]
		}
		for j := 0; j < len(strs[0]); j++ {
			if strs[0][j] != strs[i][j] {
				strs[0] = strs[0][:j]
				j = len(strs[0])
			}
		}
	}
	return strs[0]
}

func main() {
	longestCommonPrefix([]string{"flow", "flsdf", "floer"})
}
