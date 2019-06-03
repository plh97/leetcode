package findComplement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans int
}{
	{
		4,
		3,
	},
	{
		5,
		2,
	},
	{
		1,
		0,
	},
	{
		2,
		1,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findComplement(tc.N1), "输入:%v", tc)
	}
}
