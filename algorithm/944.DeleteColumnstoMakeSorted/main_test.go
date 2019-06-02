package minDeletionSize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []string
	ans int
}{
	{
		[]string{"cba", "daf", "ghi"},
		1,
	},
	{
		[]string{"rrjk", "furt", "guzm"},
		2,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, minDeletionSize(tc.N1), "输入:%v", tc)
	}
}
