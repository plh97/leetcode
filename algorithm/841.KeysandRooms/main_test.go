package canVisitAllRooms

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
			[]int{1},
			[]int{2},
			[]int{3},
			[]int{},
		},
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, canVisitAllRooms(tc.N1), "输入:%v", tc)
	}
}
