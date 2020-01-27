package main

import (
	"fmt"
	"os"
)

func main() {
	//CommandLine()
	//LocalVariable()
	GlobalVariable()
}

func CommandLine() {
	// 接受用户传递的参数，都是以字符串方式传递
	list := os.Args
	n := len(list)
	for i, data := range list {
		fmt.Printf("i = %v, data = %v, n = %v\n", i, data, n)
	}
}

func LocalVariable() {
	// 定义在 {} 里面的变量就是局部变量，只能在 {} 里面有效
	// 作用域，变量作用的范围
	if flag := 3; flag == 3 {
		fmt.Println("flage = ", flag)
	}
	// 外面就没有 flag 了
}

// 定义在函数外部的变量是全局变量
func test() {
	fmt.Println("a = ", a)
}

var a = 10

func GlobalVariable() {
	fmt.Println("全局的 a = ", a)
	// 1、 不同作用域，允许定义同名变量
	// 2、 使用变量的原则，就近原则
	a := 20
	fmt.Println("函数里的 a = ", a)
	{
		a := 1.1
		fmt.Println("函数里的 括号里的 a = ", a)
	}
	test()
}
