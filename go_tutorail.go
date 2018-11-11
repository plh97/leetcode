package main

import(
  "fmt"
  "math/cmplx"
)

func add( x, y int) int {
  return x+y
}

// 参数数据类型,     返回数值数据类型,可返回多个数据
func swap(x,y string) (string, string) {
  return y , x
}

// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
// 1.数据sum整数型参数
// 2.返回整数型 x,y 两个参数
func split(sum int) (x,y int) {
  x = sum * 4 / 9
  y = sum - 4
  return
}


// 变量也可以赋值类型,  赋值默认值
var i,j,k int = 1,3,5
// 也可以不赋值, 默认就是0
// var i,j,k int


func varity() {
  var i, j int  = 1,2
  // 短变量声明
  k := 3
  c, python,java := true,false,"no!"
  fmt.Println(i, j, k, c, python, java)
}



var (
  ToBe    bool        = false
  MaxInt  uint64      = 1 << 64 - 1
  z       complex128  = cmplx.Sqrt(-5 - 12i)
)
func fundamtype() {

}


func main() {
  fundamtype()
  varity()
}