package heightChecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans int
}{
	{
		[]int{1,1,4,2,1,3},
		3,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, heightChecker(tc.N1), "输入:%v", tc)
	}
}
