package combinationSum

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
	N1  []int
	N2  int
	ans [][]int
}{
	{
		[]int{2, 3, 6, 7},
		7,
		[][]int{
			[]int{2, 2, 3},
			[]int{7},
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, combinationSum(tc.N1, tc.N2), "输入:%v", tc)
	}
}
