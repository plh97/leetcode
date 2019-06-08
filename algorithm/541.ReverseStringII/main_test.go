package hammingWeight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  uint32
	ans int
}{
	{
		00000000000000000000000000001011,
		3,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, hammingWeight(tc.N1), "输入:%v", tc)
	}
}
