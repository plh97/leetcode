package removeDuplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	ans string
}{
	{
		"abbaca",
		"ca",
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, removeDuplicates(tc.N1), "输入:%v", tc)
	}
}
