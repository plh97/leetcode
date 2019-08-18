package carPooling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	N2  int
	ans bool
}{
	// {
	// 	[][]int{
	// 		[]int{2, 1, 5},
	// 		[]int{3, 3, 7},
	// 	},
	// 	4,
	// 	false,
	// },
	// {
	// 	[][]int{
	// 		[]int{7, 5, 6},
	// 		[]int{6, 7, 8},
	// 		[]int{10, 1, 6},
	// 	},
	// 	16,
	// 	false,
	// },
	{
		[][]int{
			[]int{2, 1, 5},
			[]int{3, 5, 7},
		},
		3,
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, carPooling(tc.N1, tc.N2), "输入:%v", tc)
	}
}
