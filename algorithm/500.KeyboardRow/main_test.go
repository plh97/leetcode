package findWords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []string
	ans []string
}{
	{
		[]string{"Hello", "Alaska", "Dad", "Peace"},
		[]string{"Alaska", "Dad"},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findWords(tc.N1), "输入:%v", tc)
	}
}
