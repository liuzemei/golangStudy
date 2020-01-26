package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("%t\n", flag) // true
	// bool 类型不能转换为 int
	// int 类型也不能转换为 bool

	var ch byte
	ch = 'a' // 字符型本质上就是整型
	var t int
	t = int(ch)
	fmt.Println("t = ", t)
}