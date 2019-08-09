package myPow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

var tcs = []struct {
	N1  float64
	N2  int
	ans float64
}{
	{
		2.0,
		-2,
		0.25,
	},
	{
		2.0,
		0,
		1,
	},
	{
		2.0,
		10,
		1024.0,
	},
	{
		2.00000,
		-2147483648,
		0,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, myPow(tc.N1, tc.N2), "输入:%v", tc)
	}
}
