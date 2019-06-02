package fizzBuzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		mod3 := (i%3 == 0)
		mod5 := (i%5 == 0)
		if mod3 && mod5 {
			res = append(res, "FizzBuzz")
		} else if mod3 {
			res = append(res, "Fizz")
		} else if mod5 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
