package dominantIndex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans int
}{
	{
		[]int{3, 6, 1, 0},
		1,
	},
	{
		[]int{1, 2, 3, 4},
		-1,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, dominantIndex(tc.N1), "输入:%v", tc)
	}
}
