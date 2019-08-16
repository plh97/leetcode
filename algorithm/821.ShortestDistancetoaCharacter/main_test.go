package shortestToChar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	N2  byte
	ans []int
}{
	{
		"loveleetcode",
		'e',
		[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, shortestToChar(tc.N1, tc.N2), "输入:%v", tc)
	}
}
