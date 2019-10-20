package missingnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans int
}{
	{
		[]int{5, 7, 11, 13},
		9,
	},
	{
		[]int{15, 13, 12},
		14,
	},
	{
		[]int{100, 300, 400},
		200,
	},
}

func Test(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, missingNumber(tc.N1), "输入:%v", tc)
	}
}
