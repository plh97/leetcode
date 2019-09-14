package countAndSay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans string
}{
	// {
	// 	1,
	// 	"1",
	// },
	// {
	// 	2,
	// 	"11",
	// },
	// {
	// 	3,
	// 	"21",
	// },
	// {
	// 	4,
	// 	"1211",
	// },
	// {
	// 	5,
	// 	"111221",
	// },
	// {
	// 	6,
	// 	"312211",
	// },
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, countAndSay(tc.N1), "输入:%v", tc)
	}
}
