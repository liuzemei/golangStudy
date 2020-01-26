package main

import "fmt"

func main() {
	// 变量，程序运行期间，可以改变的量

	// 1、声明格式 var 变量名 类型
	// 2、只是声明没有初始化的变量，默认值是 零值
	var a int
	var b, c int
	fmt.Println(a, b, c)

	// 变量的初始化，声明变量时，同时赋值
	var d int = 10 // 初始化，声明变量时，同时赋值（一步到位）
	d = 20
	fmt.Println("d =", d)

	// 3、自动推导类型，必须初始化，通过初始化的值确定类型(常用)
	// :=, 自动推导类型，先声明变量 b，再给 b 赋值为 20
	e := 30
	fmt.Printf("c type is %T\n", e)
}
