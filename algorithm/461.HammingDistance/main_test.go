package hammingDistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	N2  int
	ans int
}{
	{
		1,
		4,
		2,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, hammingDistance(tc.N1, tc.N2), "输入:%v", tc)
	}
}
