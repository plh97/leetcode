## [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)

###### 2019/02/22

### 题目💗
正则表达式匹配,很难但是必考题,一定要会,采用**消项**的思路


> #### 思路
> 采用消项的思路


### TODO
- [x] 测试用例通过
- [ ] 双层for循环遍历一般都很慢,应该尽量避免写O(m*n)代码

### 感想
做了这道题我才真正理解了正则匹配,什么叫做正则匹配中的**回溯算法**.这才堪堪入门算法,初窥.我有点懂了正则里面的回溯匹配了.mmp,其核心不正是这种匹配消耗策略吗,

### 😂有待纠正的错误
- [ ] 自己乱写,根本写不粗来...
  看到网上很多评论 
  ```
  easy dp solution, 也就是这题利用动态规划的思想,很容已解决...
  ```
  而我的答案,勉强通过了7/8个测试用例,
  ```go
  package isMatch

  import (
    "fmt"
    "strings"
  )

  type description struct {
    val   string
    index int
  }

  // 包括`.`检测目标字符串是否匹配
  func isMatchIncludeDot(s, p string) bool {
    if len(p) > len(s) {
      return false
    }
    for i := range p {
      _p := p[i : i+1]
      if _p != "." && s[i:i+1] != _p {
        return false
      }
    }
    return true
  }

  // 字符串匹配消耗
  // s: aa -> p: a  匹配消耗策略: 贪婪匹配
  func consume(s, p string) string {
    fmt.Println(s, p)
    if p == "*" {
      return s
    }
    if strings.Index(p, "*") > -1 {
      // 整个匹配无法消耗,尝试分割匹配
      // ppp pp* 优先采用分割匹配策略
      // if len(p)>1 && len(s)>len(p)-1 {
      // 	return consume
      // 	// for
      // }
      // 无限匹配策略
      ss := s
      for isMatchIncludeDot(ss, p[:len(p)-1]) == true {
        ss = ss[len(p)-1:]
      }
      return ss
    } else {
      if len(s) < len(p) {
        return s
      }
      // 对比两个字符串是否相等, 遇到 `.` 通配符
      for i := range p {
        _p := p[i : i+1]
        if _p != "." && s[i:i+1] != _p {
          return s
        }
      }
      return s[len(p):]
    }
  }

  // 返回当前匹配需要消除的字符串
  func getMatch(s string) description {
    // mississippi
    // mis*is*ip*.
    // 匹配关键字符:
    // 匹配策略,
    // 字符串匹配不做消耗,
    // 字符串* 消耗一次
    indexStar := strings.Index(s, "*")
    if indexStar == -1 {
      return description{
        val:   s,
        index: len(s),
      }
    }
    // 匹配字符串模式
    return description{
      // 值
      val: s[:indexStar+1],
      // 模式
      index: indexStar + 1,
    }
  }

  func isMatch(s string, p string) bool {
    // 采用成功匹配就销项的策略
    //
    for len(p) > 0 && len(s) > 0 {
      description := getMatch(p)
      p = p[description.index:]
      s = consume(s, description.val)

      fmt.Printf("要被消耗的字符串: %s\t 剩余字符串: %s\n", description.val, p)
      fmt.Println("===========================================================", s)
    }
    fmt.Println("===========================================================")
    return len(s) == 0
  }
  ```
  自己写根本不可能写出来啊, 除非我创造一个动态规划,这可能么??