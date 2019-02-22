package isMatch

import (
	"fmt"
	"strings"
)

type times struct {
	val  string
	time int
}

type description struct {
	val   string
	mode  string
	index int
}

// 返回当前匹配需要消除的字符串
func getMatch(s string) description {
	indexStar := strings.Index(s, "*")
	if indexStar == -1 {
		return description{
			val:   s,
			mode:  "",
			index: len(s),
		}
	}
	return description{
		val:   s[:indexStar+1],
		mode:  "*",
		index: indexStar + 1,
	}
}

// 字符串匹配消耗
// s: aa -> p: a  匹配消耗策略: 贪婪匹配
func consume(s, p string) string {
	fmt.Println(s, p)
	if strings.Index(p, "*") > -1 {
		// 无限匹配策略
	} else if strings.Index(p, ".") > -1 {
		// 
	}
	return s
}

func isMatch(s string, p string) bool {
	_p := p
	i := 20
	// 采用成功匹配就销项的策略
	for len(_p) > 0 && i > 0 {
		i--
		description := getMatch(_p)
		_p = _p[description.index:]
		s = consume(s, description.val)

		fmt.Printf("要被消耗的字符串: %s\t 匹配模式:%s\t 剩余字符串: %s\n", description.val, description.mode, _p)
		fmt.Println("===========================================================")
	}
	fmt.Println("===========================================================")
	return len(s) == 0
}
