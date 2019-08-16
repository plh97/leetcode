package countBits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans []int
}{
	{
		2,
		[]int{0, 1, 1},
	},
	{
		5,
		[]int{0, 1, 1, 2, 1, 2},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, countBits(tc.N1), "输入:%v", tc)
	}
}
