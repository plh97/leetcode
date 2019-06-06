package convertToBase7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans string
}{
	{
		100,
		"202",
	},
	{
		-7,
		"-10",
	},
	{
		0,
		"0",
	},
	{
		-8,
		"-1",
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, convertToBase7(tc.N1), "输入:%v", tc)
	}
}
