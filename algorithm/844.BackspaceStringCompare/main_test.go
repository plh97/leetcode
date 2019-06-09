package backspaceCompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	N2  string
	ans bool
}{
	{
		"ab#c",
		"ad#c",
		true,
	},
	{
		"ab##",
		"c#d#",
		true,
	},
	{
		"a##",
		"c#d#",
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, backspaceCompare(tc.N1, tc.N2), "输入:%v", tc)
	}
}
