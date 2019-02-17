
[![LeetCode 排名](https://img.shields.io/badge/pengliheng-6-blue.svg)](https://leetcode.com/pengliheng/)
[![codecov](https://codecov.io/gh/pengliheng/teetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/pengliheng/leetcode)
[![Build Status](https://www.travis-ci.org/pengliheng/leetcode.svg?branch=master)](https://www.travis-ci.org/aQuaYi/LeetCode-in-Go)


### Go 语言是一个非常强大的语言 很棒!

Go 语言接受了函数式编程的一些想法，支持匿名函数与闭包。再如，Go 语言接受了以 Erlang 语言为代表的面向消息编程思想，支持 goroutine 和通道，并推荐使用消息而不是共享内存来进行并发编程。总体来说，Go 语言是一个非常现代化的语言，精小但非常强大。



### Go 语言最主要的特性：
- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

### Go有什么优势
- **可直接编译成机器码** 不依赖其他库，glibc的版本有一定要求，部署就是扔一个文件上去就完成了。
  
- **静态类型语言** 但是有动态语言的感觉，静态类型的语言就是可以在编译的时候检查出来隐藏的大多数问题，动态语言的感觉就是有很多的包可以使用，写起来的效率很高。
  
- **语言层面支持并发** 这个就是Go最大的特色，天生的支持并发，我曾经说过一句话，天生的基因和整容是有区别的，大家一样美丽，但是你喜欢整容的还是天生基因的美丽呢？Go就是基因里面支持的并发，可以充分的利用多核，很容易的使用并发。
  
- **内置runtime** 支持垃圾回收，这属于动态语言的特性之一吧，虽然目前来说GC不算完美，但是足以应付我们所能遇到的大多数情况，特别是Go1.1之后的GC。

- **简单易学** Go语言的作者都有C的基因，那么Go自然而然就有了C的基因，那么Go关键字是25个，但是表达能力很强大，几乎支持大多数你在其他语言见过的特性：继承、重载、对象等。

- **丰富的标准库** Go目前已经内置了大量的库，特别是网络库非常强大，我最爱的也是这部分。

- **内置强大的工具** Go语言里面内置了很多工具链，最好的应该是gofmt工具，自动化格式化代码，能够让团队review变得如此的简单，代码格式一模一样，想不一样都很困难。

- **跨平台编译** 如果你写的Go代码不包含cgo，那么就可以做到window系统编译linux的应用，如何做到的呢？Go引用了plan9的代码，这就是不依赖系统的信息。

- **内嵌C支持** 前面说了作者是C的作者，所以Go里面也可以直接包含c代码，利用现有的丰富的C库。



### Go适合用来做什么
- **服务器编程** 以前你如果使用C或者C++做的那些事情，用Go来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。

- **分布式系统** 数据库代理器等

- **网络编程** 这一块目前应用最广，包括Web应用、API应用、下载应用、

- **内存数据库** 前一段时间google开发的groupcache，couchbase的部分组建

- **云平台** 目前国外很多云平台在采用Go开发，CloudFoundy的部分组建，前VMare的技术总监自己出来搞的apcera云平台。


### Go 语言的基础

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

### Go 的 25 个关键字

|  break   |   default   |  func  | interface | select |
| :------: | :---------: | :----: | :-------: | :----: |
|   case   |    defer    |   go   |    map    | struct |
|   chan   |    else     |  goto  |  package  | switch |
|  const   | fallthrough |   if   |   range   |  type  |
| continue |     for     | import |  return   |

### Go 的基本类型
|   bool    |   布尔值    |                       |        |        |         |
| :-------: | :---------: | :-------------------: | :----: | :----: | :-----: |
|  string   |   字符串    |
|    int    |    int8     |         int16         | int32  | int64  |
|   uint    |    uint8    |        uint16         | uint32 | uint64 | uintptr |
|  float32  |   float64   |      小数/浮点数      |
| complex64 | complete128 |  这你妈是复数啊,擦!   |
|   byte    |  uint8别名  |
|   rune    |  int32别名  | 表示一个 Unicode 码点 |


### fmt 基础库
- **Print** 
  it will print number variables, and will not include a line break at the end.
  它会打印数字变量,但是将不会包括一行的尾部断点.

- **Printf** 
  it will not print number variables, and will not include a line break at the end.
  它将不会打印数字变量, 并且将不会包括一行的尾部断点.

- **Println** 
  it will print number variables, and will include a line break at the end.
  它将会打印数字变量,并且将会包括一行尾部的断行.



### bug free
为什么我的代码总被同事吐槽bug多呢?
浅层的考虑是,考虑的不够周全,
深层的考虑是,写代码的习惯不好.
例如:
- `leetcode`做题不能一次性通过.老是修改,因为在计算机里面修改文字是没有成本的
- 但是一个工程师思维不严谨,造成的问题是奔溃性的,无法修复的.
- 这样持续下去会造成什么样的问题呢?
- 写代码很随意,没有经过深思熟虑.写错了直接编译一次嘛.
- 但是很多bug是无法通过编译能够发现的.尤其是JavaScript这种动态语言,
- 所以一定要养成好习惯.尤其是在写JavaScript这种动态语言的时候.


### 我的观念
程序员的职场核心竞争力是coding的能力,这个是硬指标
软指标 忠诚度, 做事靠谱程度等