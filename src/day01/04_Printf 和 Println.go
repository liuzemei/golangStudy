package main

import "fmt"

func main() {
	a := 10
	// 一段一段处理，自动加换行
	fmt.Println("a = ", a)

	// 格式化输出， 把 a 的内容放在 %d 的位置
	fmt.Printf("a = %d\n", a)
}
