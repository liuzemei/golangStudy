package main

import "fmt"

func main() {
	//AnonymousFunc()
	Closure()
}

func AnonymousFunc() {

	a := 10
	str := "mike"

	// 匿名函数，没有函数名字，函数定义，还没有调用
	f1 := func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}
	// 函数执行
	f1()

	// 给函数类型起别名
	type FuncType func() //函数没有参数，没有返回值
	// 声明变量
	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数，同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}()

	// 带参数的匿名函数
	func(i, j int) {
		fmt.Printf("i = %d, str = %d \n", i, j)
	}(1, 2)

	// 带参数带返回值的匿名函数
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Printf("x = %d, y = %d \n", x, y)
}

func Closure() {
	a := 10
	str := "mike"

	func() {
		// 1. 闭包是以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部: a = %d, str = %s\n", a, str)
	}()
	fmt.Printf("内部: a = %d, str = %s\n", a, str)

	// 返回值为一个匿名函数，通过 f 来调用返回的匿名函数，f 来调用闭包函数
	// 它不关心这些捕获了的变量和常量是否已经超出了作用域
	// 所以只要有闭包还在使用它，这些变量就还会存在
	f := testFunc()
	b := f()
	fmt.Println("b = ", b) // 1
	b = f()
	fmt.Println("b = ", b) // 4
	b = f()
	fmt.Println("b = ", b) // 9
}

func testFunc() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
