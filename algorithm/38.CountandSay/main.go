package countAndSay

import (
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	for n > 1 {
		n--
		temp := ""
		j := 0
		for j < len(res) {

			count := 1
			e := res[j]

			for j < len(res)-1 && res[j] == res[j+1] {
				count++
				j++
			}
			j++
			temp += strconv.Itoa(count)
			temp += string(e)

		}
		res = temp
	}
	return res
}
