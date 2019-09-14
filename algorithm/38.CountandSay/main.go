package countAndSay

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	for n > 1 {
		fmt.Println("=============times ============", n, res)
		n--
		temp := ""
		for len(res) > 0 {
			times := 0
			for len(res) > 0 && res[0] == '1' {
				times++
				res = res[1:]
			}
			if times == 1 {
				if len(res) == 0 {
					temp += "11"
					continue
				} else {
					temp += "1"
				}
			} else if times == 0 {
				temp += "1"
				temp += res[0:1]
			} else if times > 1 {
				temp += strconv.Itoa(times)
				temp += "1"
			}
			if len(res) > 0 {
				res = res[1:]
			}
		}
		res = temp
	}
	return res
}
