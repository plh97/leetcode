package numJewelsInStones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	N2  string
	ans int
}{
	{
		"aA",
		"aAAbbbb",
		3,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, numJewelsInStones(tc.N1, tc.N2), "输入:%v", tc)
	}
}
