package lastStoneWeight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans int
}{
	{
		[]int{2, 7, 4, 1, 8, 1},
		1,
	},
	{
		[]int{0},
		0,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, lastStoneWeight(tc.N1), "输入:%v", tc)
	}
}
