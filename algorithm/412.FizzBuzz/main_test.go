package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans []string
}{
	{
		15,
		[]string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, fizzBuzz(tc.N1), "输入:%v", tc)
	}
}
