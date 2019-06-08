package reverseStr

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
		"abcdefg",
		5,
		"edcbafg",
	},
	{
		"hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl",
		39,
		"fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi",
	},
	{
		"abcdefg",
		8,
		"gfedcba",
	},
	{
		"abcdefg",
		2,
		"bacdfeg",
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, reverseStr(tc.N1, tc.N2), "输入:%v", tc)
	}
}
