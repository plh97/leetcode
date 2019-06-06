package randomMoney

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Map map[int]int = make(map[int]int, 20)

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for i := 0; i < 51111; i++ {
		sum, res := randomMoney(100)
		for _, e := range res {
			if _, ok := Map[e]; !ok {
				Map[e] = 0
			} else {
				Map[e]++
			}
			if e < 6 || e > 12 {
				ast.Equal(100, 101, "%s不在6-12内", e)
			}
		}
		ast.Equal(100, sum, "sum: %s", sum)
	}
	fmt.Println(Map)
}
