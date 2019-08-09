package dailyTemperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans []int
}{
	{
		[]int{73, 74, 75, 71, 69, 72, 76, 73},
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, dailyTemperatures(tc.N1), "输入:%v", tc)
	}
}
