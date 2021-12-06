// pointer 指针
package main

import "fmt"

func main() {
	// 声明变量 x 赋值为 1
	x := 1
	// 声明变量 p 赋值为 x 的变量的内存地址
	p := &x
	fmt.Print(*p)
	// 修改 p 对应内存地址的值
	*p = 2
	fmt.Print(x)
}
