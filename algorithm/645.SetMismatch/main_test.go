package findErrorNums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans []int
}{
	{
		[]int{1, 2, 2, 4},
		[]int{2, 3},
	},
	{
		[]int{2, 2},
		[]int{2, 1},
	},
	{
		[]int{1, 1},
		[]int{1, 2},
	},
	{
		[]int{3, 2, 2},
		[]int{2, 1},
	},
	{
		[]int{3, 3, 1},
		[]int{3, 2},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findErrorNums(tc.N1), "输入:%v", tc)
	}
}
