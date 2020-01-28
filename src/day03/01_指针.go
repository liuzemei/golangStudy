package main

import "fmt"

func main() {
	//pointer()
	//newStudy()
	pointerUse()
}

func pointer() {
	// 每个变量有 2 层含义：变量的内存 和 变量的地址
	var a int = 10
	fmt.Printf("a = %d\n", a)   //变量的内存
	fmt.Printf("&a = %p\n", &a) //变量的地址
	// 保存某个变量的地址，需要指针类型 *int 保存 int 的地址， **int 保存 *int 地址
	// 定义一个变量 p，类型为 *int
	var p *int
	p = &a
	*p = 666
	fmt.Printf("p = %p\n", p) //
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&p = %p\n", &p) //
}

func newStudy() {
	//a := 10
	var p *int
	// new() 返回的是一个指针，指向特定数据类型的起始位置。
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)

	q := new(int)
	*q = 777
	fmt.Println("*q = ", *q)
}

func pointerUse() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(p1 *int, p2 *int) {
	*p1, *p2 = *p2, *p1
}
