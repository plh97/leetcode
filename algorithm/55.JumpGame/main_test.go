package canJump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans bool
}{
	{
		[]int{2, 3, 1, 1, 4},
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, canJump(tc.N1), "输入:%v", tc)
	}
}
