package isPowerOfFour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans bool
}{
	{
		4,
		true,
	},
	{
		16,
		true,
	},
	{
		15,
		false,
	},
	{
		1,
		true,
	},
	{
		8,
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isPowerOfFour(tc.N1), "输入:%v", tc)
	}
}
