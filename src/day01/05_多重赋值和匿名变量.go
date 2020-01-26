package main

import "fmt"

func main() {
	// 多重变量的值
	a, b := 10, 20
	// 快捷交换
	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)

	tmp, _ := a, b
	fmt.Println("tmp = ", tmp)

	// _匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势
	c, d, _ := test()
	fmt.Printf("c = %d, d = %d\n", c, d)
}

func test() (a, b, c int) {
	return 1, 2, 3
}
