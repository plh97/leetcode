package convert

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	var res int64 = 0
	sign := 1
	fmt.Println("before reverse: ", x)
	if x < 0 {
		sign = -1
	}
	// if x > 0 {
	// 	ss := strings.Split(strconv.Itoa(x), "")
	// 	for i := (len(ss) - 1) / 2; i >= 0; i-- {
	// 		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	// 	}
	// 	res, _ = strconv.ParseInt(strings.Join(ss, ""), 10, 32)
	// } else {
	ss := strings.Split(strconv.Itoa(-x), "")
	for i := (len(ss) - 1) / 2; i >= 0; i-- {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
	res, _ = strconv.ParseInt(strings.Join(ss, ""), 10, 32)
	_res:=int(res)
	_res = sign * _res
	// }
	fmt.Println("after reverse: ", res)
	fmt.Println(int(res) > math.MaxInt32 || int(res) < math.MinInt32)
	fmt.Println("==============")
	if int(res) > math.MaxInt32 || int(res) < math.MinInt32 {
		res = 0
	}
	return int(res)
}
