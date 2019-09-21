package exist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]byte
	N2  string
	ans bool
}{
	{
		[][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		"ABCCED",
		true,
	},
	{
		[][]byte{
			[]byte{'a', 'a'},
		},
		"aa",
		true,
	},
	{
		[][]byte{
			[]byte{'b', 'a'},
		},
		"ab",
		true,
	},
	{
		[][]byte{
			[]byte{'a', 'a'},
		},
		"aaa",
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, exist(tc.N1, tc.N2), "输入:%v", tc)
	}
}
