package orangesRotting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans bool
}{
	{
		[][]int{
			[]int{2, 1, 1},
			[]int{1, 1, 0},
			[]int{0, 1, 1},
		},
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, orangesRotting(tc.N1), "输入:%v", tc)
	}
}
