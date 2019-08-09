package queryString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	N2  int
	ans bool
}{
	{
		"0110",
		3,
		true,
	},
	{
		"1111000101",
		5,
		true,
	},
	{
		"10010111100001110010",
		10,
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, queryString(tc.N1, tc.N2), "输入:%v", tc)
	}
}
