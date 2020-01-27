package main

import "fmt"

func main() {
	fmt.Printf("Add type is %T\n", Add) // func(int, int) int
	useFuncType()
}

func Add(a, b int) int {
	return a + b
}

// 函数也是一种数据类型，通过 type 给一个函数类型起名
// FuncType 它是一个函数类型，有两个 int 形参和 1 个 int 返回值的函数类型
type FuncType func(int, int) int

func useFuncType() {
	var fTest FuncType
	fTest = Add
	res := fTest(1, 2)
	fmt.Println("res = ", res)
}

// 回调函数，函数有一个参数是函数类型，这个函数就是回调函数
