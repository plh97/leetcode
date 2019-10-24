package getpermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	N2  int
	ans string
}{
	{
		3,
		3,
		"213",
	},
}

func Test(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, getPermutation(tc.N1,tc.N2), "输入:%v", tc)
	}
}
