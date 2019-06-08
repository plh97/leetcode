package rotateString

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
		"abcde",
		"deabc",
		true,
	},
	{
		"aa",
		"a",
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, rotateString(tc.N1, tc.N2), "输入:%v", tc)
	}
}
