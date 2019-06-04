package twoCitySchedCost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans int
}{
	{
		[][]int{
			[]int{10, 20},
			[]int{30, 200},
			[]int{400, 50},
			[]int{30, 20},
		},
		110,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, twoCitySchedCost(tc.N1), "输入:%v", tc)
	}
}
