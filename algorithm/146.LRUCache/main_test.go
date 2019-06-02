package capacity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	ast.Equal(1, cache.Get(1))
	cache.Put(3, 3) // evicts key 2
	ast.Equal(-1, cache.Get(2))
	cache.Put(4, 4) // evicts key 1
	ast.Equal(-1, cache.Get(1))
	ast.Equal(3, cache.Get(3))
	ast.Equal(4, cache.Get(4))
	cache.Put(4, 4) // evicts key 1

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
