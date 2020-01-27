package main

import "fmt"

func main() {
	// 1. 无返回值
	//MyFunc1() // 无参
	//MyFunc2(11, 12) // 有参
	//MyFunc3(1, 2, 3, 4) // 不定参
	//MyFunc4() // 不定参传递

	// 2. 有返回值
	MyFunc5()
}

// 1. 无参无返回值函数的定义
func MyFunc1() {
	a := "无参无返回值"
	fmt.Println("1: ", a)
}

// 2. 有参无返回值函数的定义
func MyFunc2(a int, b int) {
	fmt.Println("参数是：", a, b)
}

// 3. 不定参数类型
// ...type 不定参数类型
func MyFunc3(args ...int) {
	for i, data := range args {
		fmt.Printf("i[%d] = [%d]\n", i, data)
	}
}

// 注意，不定参数只能放在形参中的最后一个参数
func MyFunc3_1(a int, args ...int) {
}

// 4. 不定参数传递
func MyFunc4() {
	MyFunc4_1(1, 2, 3, 4, 5)
}

func MyFunc4_1(tmp ...int) {
	// 1. 全部形参传递给 4_2
	//MyFunc7(tmp...)

	// 2. 只想把前两个形参传递给 4_2
	MyFunc4_2(tmp[:2]...) // 0 ~ 2 （不包括 2），传递过去
	MyFunc4_2(tmp[2:]...) // 从 2 开始（包括 2），把后面所有的元素传递过去
}

func MyFunc4_2(args ...int) {
	fmt.Println("MyFunc7 :", args)
}

// 5. 无参有返回值
// 只有一个返回值
func MyFunc5() int {
	a := MyFunc5_1()
	fmt.Println("a = ", a)
	b, c := MyFunc5_2()
	fmt.Printf("b = %d, c = %d\n", b, c)
	return 666
}

// 给返回值起个名字
func MyFunc5_1() (result int) {
	result = 777
	return
}

// 有多个返回值
func MyFunc5_2() (res1 int, res2 int) {
	res1, res2 = 666, 777
	return
}
