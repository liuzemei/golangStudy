package main

import "fmt"

func main() {
	//test01()
	test02()
}

func test01() {
	a, b := 10, 20
	defer func() {
		// 变量的值会更改
		fmt.Printf("a = %d, b = %d\n", a, b)
	}()
	a, b = 111, 222
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func test02() {
	a, b := 10, 20
	defer func(a, b int) {
		// () 代表调用此匿名函数， 把参数传递过去，已经先传递参数了，只是没有调用
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(a, b)
	a, b = 111, 222
	fmt.Printf("a = %d, b = %d\n", a, b)
}
