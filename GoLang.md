
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
- **服务器编程** 以前如果使用`C`或者`C++`做的那些事情，用Go来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。
- **分布式系统** 数据库代理器等
- **网络编程** 这一块目前应用最广，包括Web应用、API应用、下载应用、
- **内存数据库** 前一段时间 google 开发的 groupCache，couchBase的部分组建
- **云平台** 目前国外很多云平台在采用`Go`开发，`CloudFoundy`的部分组建，前`VMare`的技术总监自己出来搞的apcera云平台。

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
|   name    |     名字     |                       |        |        |         |
| :-------: | :---------: | :-------------------: | :----: | :----: | :-----: |
|   bool    |   布尔值     |                       |        |        |         |
|  string   |   字符串     |
|    int    |    int8     |         int16         | int32  | int64  |
|   uint    |    uint8    |        uint16         | uint32 | uint64 | uintptr |
|  float32  |   float64   |      小数/浮点数      |
| complex64 | complete128 |         复数          |
|   byte    |  uint8别名   |
|   rune    |  int32别名   | 表示一个 Unicode 码点 |

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
