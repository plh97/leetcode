[![LeetCode 完成数](https://img.shields.io/badge/pengliheng-6-blue.svg)](https://leetcode.com/pengliheng/)
[![codeCov](https://codecov.io/gh/pengliheng/leetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/pengliheng/leetcode)
[![Build Status](https://www.travis-ci.org/pengliheng/leetcode.svg?branch=master)](https://www.travis-ci.org/pengliheng/leetcode)

### Golang 是一个非常强大的语言 很棒!
Go 语言接受了函数式编程的一些想法，支持匿名函数与闭包。再如，Go 语言接受了以 ErLang 语言为代表的面向消息编程思想，支持 goroutine 和通道，并推荐使用消息而不是共享内存来进行并发编程。总体来说，Go 语言是一个非常现代化的语言，精小但非常强大。

### Golang 最主要的特性：
- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

### Golang 适合用来做什么
- **服务器编程** 以前如果使用C或者C++做的那些事情，用Go来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。
- **分布式系统** 数据库代理器等
- **网络编程** 这一块目前应用最广，包括Web应用、API应用、下载应用、
- **内存数据库** 前一段时间 google 开发的 groupCache，couchBase的部分组建
- **云平台** 目前国外很多云平台在采用Go开发，CloudFoundy的部分组建，前VMare的技术总监自己出来搞的apcera云平台。

### Golang 语言的基础
- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

### Golang 的25个关键字
|  break   |   default   |  func  | interface | select |
| :------: | :---------: | :----: | :-------: | :----: |
|   case   |    defer    |   go   |    map    | struct |
|   chan   |    else     |  goto  |  package  | switch |
|  const   | fallthrough |   if   |   range   |  type  |
| continue |     for     | import |  return   |

### Golang 的基本类型
|   bool    |   布尔值    |                       |        |        |         |
| :-------: | :---------: | :-------------------: | :----: | :----: | :-----: |
|  string   |   字符串    |
|    int    |    int8     |         int16         | int32  | int64  |
|   uint    |    uint8    |        uint16         | uint32 | uint64 | uintptr |
|  float32  |   float64   |      小数/浮点数      |
| complex64 | complete128 |         复数          |
|   byte    |  uint8别名  |
|   rune    |  int32别名  | 表示一个 Unicode 码点 |

### fmt 基础库
- **Print** 
  it will print number variables, and will not include a line break at the end.
  它会打印数字变量,但是将不会包括一行的尾部断点.
- **Printf** 
  it will not print number variables, and will not include a line break at the end.
  它将不会打印数字变量, 并且将不会包括一行的尾部断点.
- **Println** 
  it will print number variables, and will include a line break at the end.
  它将会打印数字变量,并且将会包括一行尾部的断行.

### Bug Free
为什么我的代码总被同事吐槽bug多呢?
浅层的考虑是,考虑的不够周全,
深层的考虑是,写代码的习惯不好.
例如:
- `LeetCode`做题不能一次性通过.老是修改,因为在计算机里面修改文字是没有成本的
- 但是一个工程师思维不严谨,造成的问题是奔溃性的,无法修复的.
- 这样持续下去会造成什么样的问题呢?
- 写代码很随意,没有经过深思熟虑.写错了直接编译一次嘛.
- 但是很多bug是无法通过编译能够发现的.尤其是JavaScript这种动态语言,
- 所以一定要养成好习惯.尤其是在写JavaScript这种动态语言的时候.

### 我的观念
- *硬指标* 程序员的职场核心竞争力是coding的能力, 这个靠刷题可以极大地提升coding能力
- *软指标* 忠诚度, 做事靠谱程度等


### 采访一位刷了1000道leetcode的[大佬](https://github.com/aQuaYi/LeetCode-in-Go)
我还是觉得得刷题
```
ali peng, [16.02.19 15:34]
大佬您好, 看到您用golang的刷题记录, 感觉真的很棒,特地来电报膜拜一下

aQua Yi, [18.02.19 20:39]
谢谢，我只是把我做过的题目记录了下来而已。

ali peng, [18.02.19 20:40]
🙃

ali peng, [18.02.19 20:41]
大佬 能否给刚入职场一年的新人点建议,, 刷题有作用么, 对于golang后端开发来说

aQua Yi, [18.02.19 21:46]
哈哈，我是一名机械工程师，还没有入职呢。学习编程是业余爱好。

aQua Yi, [18.02.19 21:47]
不过，刷题确实很锻炼能力，我现在看我早期的刷题记录，都觉得很烂。

ali peng, [18.02.19 22:49]
🙃谢谢反馈,

ali peng, [18.02.19 22:52]
😂感觉大佬您的编程能力比我身边的同事都厉害

ali peng, [18.02.19 22:53]
让我感觉我们写代码就光只会调用几个函数库了
```