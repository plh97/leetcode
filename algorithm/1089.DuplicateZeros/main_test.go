package duplicateZeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans []int
}{
	{
		[]int{1, 0, 2, 3, 0, 4, 5, 0},
		[]int{1, 0, 0, 2, 3, 0, 0, 4},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, duplicateZeros(tc.N1), "输入:%v", tc)
	}
}
