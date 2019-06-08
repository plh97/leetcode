package licenseKeyFormatting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	N2  int
	ans string
}{
	{
		"a-a-a-a-",
		1,
		"A-A-A-A",
	},
	{
		"---",
		3,
		"",
	},
	{
		"5F3Z-2e-9-w",
		4,
		"5F3Z-2E9W",
	},
	{
		"2-5g-3-J",
		2,
		"2-5G-3J",
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, licenseKeyFormatting(tc.N1, tc.N2), "输入:%v", tc)
	}
}
